// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ValidateApplicationResponseBody
	GetCode() *int32
	SetData(v string) *ValidateApplicationResponseBody
	GetData() *string
	SetMessage(v string) *ValidateApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ValidateApplicationResponseBody
	GetRequestId() *string
}

type ValidateApplicationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the application.
	//
	// example:
	//
	// 123
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ValidateApplicationResponseBody) GetData() *string {
	return s.Data
}

func (s *ValidateApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ValidateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateApplicationResponseBody) SetCode(v int32) *ValidateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateApplicationResponseBody) SetData(v string) *ValidateApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *ValidateApplicationResponseBody) SetMessage(v string) *ValidateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateApplicationResponseBody) SetRequestId(v string) *ValidateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
