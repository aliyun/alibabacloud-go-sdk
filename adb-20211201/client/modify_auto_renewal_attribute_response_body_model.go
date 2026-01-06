// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRenewalAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAutoRenewalAttributeResponseBody
	GetRequestId() *string
}

type ModifyAutoRenewalAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 421794A3-72A5-5D27-9E8B-A75A4C503E17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoRenewalAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRenewalAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewalAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoRenewalAttributeResponseBody) SetRequestId(v string) *ModifyAutoRenewalAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoRenewalAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
