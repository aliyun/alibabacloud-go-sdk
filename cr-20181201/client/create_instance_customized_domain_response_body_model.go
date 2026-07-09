// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceCustomizedDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceCustomizedDomainResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateInstanceCustomizedDomainResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateInstanceCustomizedDomainResponseBody
	GetRequestId() *string
}

type CreateInstanceCustomizedDomainResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1FA54F8C-8849-57F9-8069-F5F15EE82BE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceCustomizedDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceCustomizedDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceCustomizedDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceCustomizedDomainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateInstanceCustomizedDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceCustomizedDomainResponseBody) SetCode(v string) *CreateInstanceCustomizedDomainResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceCustomizedDomainResponseBody) SetIsSuccess(v bool) *CreateInstanceCustomizedDomainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateInstanceCustomizedDomainResponseBody) SetRequestId(v string) *CreateInstanceCustomizedDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceCustomizedDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
