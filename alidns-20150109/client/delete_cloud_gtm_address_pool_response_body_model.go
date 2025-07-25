// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolId(v string) *DeleteCloudGtmAddressPoolResponseBody
	GetAddressPoolId() *string
	SetRequestId(v string) *DeleteCloudGtmAddressPoolResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCloudGtmAddressPoolResponseBody
	GetSuccess() *bool
}

type DeleteCloudGtmAddressPoolResponseBody struct {
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89528023225442**16
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCloudGtmAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmAddressPoolResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *DeleteCloudGtmAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudGtmAddressPoolResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCloudGtmAddressPoolResponseBody) SetAddressPoolId(v string) *DeleteCloudGtmAddressPoolResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *DeleteCloudGtmAddressPoolResponseBody) SetRequestId(v string) *DeleteCloudGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudGtmAddressPoolResponseBody) SetSuccess(v bool) *DeleteCloudGtmAddressPoolResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCloudGtmAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
