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
	// The unique ID of the command request. The ID is valid for one day by default. In asynchronous scenarios, if you need the result, query it promptly within the validity period, preferably before the command times out.
	//
	// example:
	//
	// cmd-81de027b66e442e99c1e0e09a16a0be5
	CmdId *string `json:"CmdId,omitempty" xml:"CmdId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the command response.
	//
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
