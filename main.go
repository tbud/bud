package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	imageName := "alpine"
	containerName := "bud_" + imageName

	//resp, err := cli.ContainerCreate(ctx, &container.Config{
	//	Image: imageName,
	//	Tty:   true,
	//	//Cmd:   []string{"/bin/sh"},
	//}, nil, nil, containerName)
	//if err != nil {
	//	if strings.HasPrefix(err.Error(), "Error: No such image:") {
	//		reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	//		if err != nil {
	//			panic(err)
	//		}
	//		io.Copy(os.Stdout, reader)
	//
	//		resp, err = cli.ContainerCreate(ctx, &container.Config{
	//			Image: "alpine",
	//			Cmd:   []string{"echo", "hello world"},
	//		}, nil, nil, "")
	//
	//		if err != nil {
	//			panic(err)
	//		}
	//	} else {
	//		panic(err)
	//	}
	//}
	//
	//if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
	//	panic(err)
	//}

	resp, err := cli.ContainerExecCreate(ctx, containerName, types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{"echo", "Hello World!"},
	})

	if err != nil {
		panic(err)
	}

	response, err := cli.ContainerExecAttach(ctx, resp.ID, types.ExecStartCheck{})
	if err != nil {
		panic(err)
	}

	err = cli.ContainerExecStart(ctx, resp.ID, types.ExecStartCheck{})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, response.Reader)

	//inspect, err := cli.ContainerExecInspect(ctx, resp.ID)
	//if err != nil {
	//	panic(err)
	//}
	//
	//options := types.ContainerLogsOptions{ShowStdout: true}
	//// Replace this ID with a container that really exists
	//out, err := cli.ContainerLogs(ctx, inspect.ContainerID, options)
	//if err != nil {
	//	panic(err)
	//}
	//
	//io.Copy(os.Stdout, out)
	//
	//println(inspect.ExitCode)
}
