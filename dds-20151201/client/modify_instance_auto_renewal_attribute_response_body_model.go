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
	// The ID of the request.
	//
	// example:
	//
	// EFD65226-08CC-4C4D-B6A4-CB3C382F67B0
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
