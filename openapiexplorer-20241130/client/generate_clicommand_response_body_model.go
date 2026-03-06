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
	SetRequestId(v string) *GenerateCLICommandResponseBody
	GetRequestId() *string
}

type GenerateCLICommandResponseBody struct {
	// The CLI instruction.
	//
	// example:
	//
	// aliyun ecs DescribeRegions --ResourceType instance
	Cli *string `json:"cli,omitempty" xml:"cli,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A707AFA8-1A4C-5B2A-A165-8436C1EA38DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *GenerateCLICommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCLICommandResponseBody) SetCli(v string) *GenerateCLICommandResponseBody {
	s.Cli = &v
	return s
}

func (s *GenerateCLICommandResponseBody) SetRequestId(v string) *GenerateCLICommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCLICommandResponseBody) Validate() error {
	return dara.Validate(s)
}
