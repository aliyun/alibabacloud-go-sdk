// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingInstanceCommandsStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetRenderingInstanceCommandsStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRenderingInstanceCommandsStatusResponseBody
	GetRequestId() *string
	SetResult(v string) *GetRenderingInstanceCommandsStatusResponseBody
	GetResult() *string
	SetStatus(v string) *GetRenderingInstanceCommandsStatusResponseBody
	GetStatus() *string
}

type GetRenderingInstanceCommandsStatusResponseBody struct {
	// example:
	//
	// conn failed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Thu Jun 27 16:06:26 CST 2024
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRenderingInstanceCommandsStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingInstanceCommandsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) GetResult() *string {
	return s.Result
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) SetMessage(v string) *GetRenderingInstanceCommandsStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) SetRequestId(v string) *GetRenderingInstanceCommandsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) SetResult(v string) *GetRenderingInstanceCommandsStatusResponseBody {
	s.Result = &v
	return s
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) SetStatus(v string) *GetRenderingInstanceCommandsStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetRenderingInstanceCommandsStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
