// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCustomerResponseBody
	GetCode() *string
	SetData(v bool) *CreateCustomerResponseBody
	GetData() *bool
	SetMessage(v string) *CreateCustomerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCustomerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCustomerResponseBody
	GetSuccess() *bool
}

type CreateCustomerResponseBody struct {
	// Code indicating whether the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data indicating whether a customer was successfully created. If it\\"s "true", the Message contains CID.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Massage indicating whether the call was successful.
	//
	// example:
	//
	// 12345
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, Alibaba Cloud will track errors with this.
	//
	// example:
	//
	// A9B725C7-3DBD-576B-AC91-F6F22AB99A77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Candidate Value: True/False, which indicates whether the current API call it self was successful. It does not guarantee the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCustomerResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateCustomerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomerResponseBody) SetCode(v string) *CreateCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomerResponseBody) SetData(v bool) *CreateCustomerResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCustomerResponseBody) SetMessage(v string) *CreateCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomerResponseBody) SetRequestId(v string) *CreateCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerResponseBody) SetSuccess(v bool) *CreateCustomerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomerResponseBody) Validate() error {
	return dara.Validate(s)
}
