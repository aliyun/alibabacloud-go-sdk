// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelTaskResponseBody
	GetCode() *string
	SetData(v bool) *CancelTaskResponseBody
	GetData() *bool
	SetMessage(v string) *CancelTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelTaskResponseBody
	GetSuccess() *bool
}

type CancelTaskResponseBody struct {
	// Request status code. A value of OK indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API was invoked successfully. Valid values:
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

func (s CancelTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelTaskResponseBody) SetCode(v string) *CancelTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelTaskResponseBody) SetData(v bool) *CancelTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelTaskResponseBody) SetMessage(v string) *CancelTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTaskResponseBody) SetSuccess(v bool) *CancelTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CancelTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
