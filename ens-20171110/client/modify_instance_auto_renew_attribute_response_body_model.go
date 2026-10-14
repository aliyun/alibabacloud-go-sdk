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
	// EA3758E0-8899-17D3-9526-5F62CF33A586
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
