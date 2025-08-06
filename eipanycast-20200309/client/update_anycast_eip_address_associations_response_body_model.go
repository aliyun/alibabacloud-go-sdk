// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAnycastEipAddressAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAnycastEipAddressAssociationsResponseBody
	GetRequestId() *string
}

type UpdateAnycastEipAddressAssociationsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAnycastEipAddressAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAnycastEipAddressAssociationsResponseBody) SetRequestId(v string) *UpdateAnycastEipAddressAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsResponseBody) Validate() error {
	return dara.Validate(s)
}
