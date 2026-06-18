// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChunkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateChunkResponseBody
	GetCode() *string
	SetData(v bool) *UpdateChunkResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateChunkResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateChunkResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateChunkResponseBody
	GetStatus() *string
	SetSuccess(v bool) *UpdateChunkResponseBody
	GetSuccess() *bool
}

type UpdateChunkResponseBody struct {
	// The error status code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned upon a successful request.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// "200"
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateChunkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChunkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChunkResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChunkResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateChunkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateChunkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChunkResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateChunkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateChunkResponseBody) SetCode(v string) *UpdateChunkResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChunkResponseBody) SetData(v bool) *UpdateChunkResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateChunkResponseBody) SetMessage(v string) *UpdateChunkResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateChunkResponseBody) SetRequestId(v string) *UpdateChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChunkResponseBody) SetStatus(v string) *UpdateChunkResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateChunkResponseBody) SetSuccess(v bool) *UpdateChunkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateChunkResponseBody) Validate() error {
	return dara.Validate(s)
}
