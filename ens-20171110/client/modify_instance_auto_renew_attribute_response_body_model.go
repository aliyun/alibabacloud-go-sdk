// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyInstanceAutoRenewAttributeResponseBody
	GetCode() *int32
	SetRequestId(v string) *ModifyInstanceAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type ModifyInstanceAutoRenewAttributeResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4DD66F05-3116-4BAA-B588-52EB2E7F431D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewAttributeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyInstanceAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceAutoRenewAttributeResponseBody) SetCode(v int32) *ModifyInstanceAutoRenewAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
