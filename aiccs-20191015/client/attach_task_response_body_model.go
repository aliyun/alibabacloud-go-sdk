// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AttachTaskResponseBody
	GetCode() *string
	SetData(v int64) *AttachTaskResponseBody
	GetData() *int64
	SetMessage(v string) *AttachTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *AttachTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AttachTaskResponseBody
	GetSuccess() *bool
}

type AttachTaskResponseBody struct {
	// The request status code. A value of OK indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The quantity of successfully appended entries.
	//
	// example:
	//
	// 10
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AttachTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *AttachTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AttachTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AttachTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AttachTaskResponseBody) SetCode(v string) *AttachTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AttachTaskResponseBody) SetData(v int64) *AttachTaskResponseBody {
	s.Data = &v
	return s
}

func (s *AttachTaskResponseBody) SetMessage(v string) *AttachTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AttachTaskResponseBody) SetRequestId(v string) *AttachTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachTaskResponseBody) SetSuccess(v bool) *AttachTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AttachTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
