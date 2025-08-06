// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnycastEipAddressSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAnycastEipAddressSpecResponseBody
	GetRequestId() *string
}

type ModifyAnycastEipAddressSpecResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAnycastEipAddressSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnycastEipAddressSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAnycastEipAddressSpecResponseBody) SetRequestId(v string) *ModifyAnycastEipAddressSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAnycastEipAddressSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
