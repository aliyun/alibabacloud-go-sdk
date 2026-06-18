// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopTaskResponseBody
	GetCode() *string
	SetData(v bool) *StopTaskResponseBody
	GetData() *bool
	SetMessage(v string) *StopTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopTaskResponseBody
	GetSuccess() *bool
}

type StopTaskResponseBody struct {
	// The request status code. A value of OK indicates that the request succeeded.
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

func (s StopTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopTaskResponseBody) SetCode(v string) *StopTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopTaskResponseBody) SetData(v bool) *StopTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopTaskResponseBody) SetMessage(v string) *StopTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopTaskResponseBody) SetRequestId(v string) *StopTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTaskResponseBody) SetSuccess(v bool) *StopTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
