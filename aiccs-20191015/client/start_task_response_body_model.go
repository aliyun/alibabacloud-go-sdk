// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartTaskResponseBody
	GetCode() *string
	SetData(v bool) *StartTaskResponseBody
	GetData() *bool
	SetMessage(v string) *StartTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartTaskResponseBody
	GetSuccess() *bool
}

type StartTaskResponseBody struct {
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

func (s StartTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartTaskResponseBody) SetCode(v string) *StartTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartTaskResponseBody) SetData(v bool) *StartTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartTaskResponseBody) SetMessage(v string) *StartTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartTaskResponseBody) SetRequestId(v string) *StartTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTaskResponseBody) SetSuccess(v bool) *StartTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
