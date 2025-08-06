// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnycastEipAddressAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAnycastEipAddressAttributeResponseBody
	GetRequestId() *string
}

type ModifyAnycastEipAddressAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAnycastEipAddressAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnycastEipAddressAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAnycastEipAddressAttributeResponseBody) SetRequestId(v string) *ModifyAnycastEipAddressAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAnycastEipAddressAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
