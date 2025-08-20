// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateChainResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *UpdateChainResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateChainResponseBody
	GetRequestId() *string
}

type UpdateChainResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 85A99B10-3926-5201-958E-C06FA47F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateChainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChainResponseBody) SetCode(v string) *UpdateChainResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChainResponseBody) SetIsSuccess(v bool) *UpdateChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateChainResponseBody) SetRequestId(v string) *UpdateChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChainResponseBody) Validate() error {
	return dara.Validate(s)
}
