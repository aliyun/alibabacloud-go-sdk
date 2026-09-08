// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMessageResponseBody
	GetCode() *string
	SetData(v bool) *ReadMessageResponseBody
	GetData() *bool
	SetMessage(v string) *ReadMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMessageResponseBody
	GetSuccess() *bool
}

type ReadMessageResponseBody struct {
	// The error code returned when the call failed. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned when the call failed.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMessageResponseBody) GetData() *bool {
	return s.Data
}

func (s *ReadMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMessageResponseBody) SetCode(v string) *ReadMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageResponseBody) SetData(v bool) *ReadMessageResponseBody {
	s.Data = &v
	return s
}

func (s *ReadMessageResponseBody) SetMessage(v string) *ReadMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageResponseBody) SetRequestId(v string) *ReadMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageResponseBody) SetSuccess(v bool) *ReadMessageResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
