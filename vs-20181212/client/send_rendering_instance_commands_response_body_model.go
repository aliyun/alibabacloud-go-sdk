// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRenderingInstanceCommandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCmdId(v string) *SendRenderingInstanceCommandsResponseBody
	GetCmdId() *string
	SetRequestId(v string) *SendRenderingInstanceCommandsResponseBody
	GetRequestId() *string
	SetResult(v string) *SendRenderingInstanceCommandsResponseBody
	GetResult() *string
}

type SendRenderingInstanceCommandsResponseBody struct {
	CmdId *string `json:"CmdId,omitempty" xml:"CmdId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Thu Jun 27 16:06:26 CST 2024
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SendRenderingInstanceCommandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendRenderingInstanceCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *SendRenderingInstanceCommandsResponseBody) GetCmdId() *string {
	return s.CmdId
}

func (s *SendRenderingInstanceCommandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendRenderingInstanceCommandsResponseBody) GetResult() *string {
	return s.Result
}

func (s *SendRenderingInstanceCommandsResponseBody) SetCmdId(v string) *SendRenderingInstanceCommandsResponseBody {
	s.CmdId = &v
	return s
}

func (s *SendRenderingInstanceCommandsResponseBody) SetRequestId(v string) *SendRenderingInstanceCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendRenderingInstanceCommandsResponseBody) SetResult(v string) *SendRenderingInstanceCommandsResponseBody {
	s.Result = &v
	return s
}

func (s *SendRenderingInstanceCommandsResponseBody) Validate() error {
	return dara.Validate(s)
}
