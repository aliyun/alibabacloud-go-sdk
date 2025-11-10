// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCLICommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCli(v string) *GenerateCLICommandResponseBody
	GetCli() *string
}

type GenerateCLICommandResponseBody struct {
	// example:
	//
	// aliyun ecs DescribeRegions --ResourceType instance
	Cli *string `json:"cli,omitempty" xml:"cli,omitempty"`
}

func (s GenerateCLICommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCLICommandResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCLICommandResponseBody) GetCli() *string {
	return s.Cli
}

func (s *GenerateCLICommandResponseBody) SetCli(v string) *GenerateCLICommandResponseBody {
	s.Cli = &v
	return s
}

func (s *GenerateCLICommandResponseBody) Validate() error {
	return dara.Validate(s)
}
