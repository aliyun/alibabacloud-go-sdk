// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRenewalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetRenewalResponseBody
	GetCode() *string
	SetMessage(v string) *SetRenewalResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetRenewalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetRenewalResponseBody
	GetSuccess() *bool
}

type SetRenewalResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetRenewalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRenewalResponseBody) GoString() string {
	return s.String()
}

func (s *SetRenewalResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetRenewalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetRenewalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRenewalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetRenewalResponseBody) SetCode(v string) *SetRenewalResponseBody {
	s.Code = &v
	return s
}

func (s *SetRenewalResponseBody) SetMessage(v string) *SetRenewalResponseBody {
	s.Message = &v
	return s
}

func (s *SetRenewalResponseBody) SetRequestId(v string) *SetRenewalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRenewalResponseBody) SetSuccess(v bool) *SetRenewalResponseBody {
	s.Success = &v
	return s
}

func (s *SetRenewalResponseBody) Validate() error {
	return dara.Validate(s)
}
