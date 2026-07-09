// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceCustomizedDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstanceCustomizedDomainResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateInstanceCustomizedDomainResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateInstanceCustomizedDomainResponseBody
	GetRequestId() *string
}

type UpdateInstanceCustomizedDomainResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 2EFAF75C-1FA7-5254-B044-E97291C170CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceCustomizedDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceCustomizedDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceCustomizedDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceCustomizedDomainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateInstanceCustomizedDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceCustomizedDomainResponseBody) SetCode(v string) *UpdateInstanceCustomizedDomainResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainResponseBody) SetIsSuccess(v bool) *UpdateInstanceCustomizedDomainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainResponseBody) SetRequestId(v string) *UpdateInstanceCustomizedDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
