// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoRenewalAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceAutoRenewalAttributeResponseBody
	GetRequestId() *string
}

type ModifyInstanceAutoRenewalAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAutoRenewalAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoRenewalAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewalAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceAutoRenewalAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAutoRenewalAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
