// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolBasicConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolBasicConfigResponseBody
	GetAddressPoolId() *string
	SetRequestId(v string) *UpdateCloudGtmAddressPoolBasicConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmAddressPoolBasicConfigResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmAddressPoolBasicConfigResponseBody struct {
	// The unique ID of the address pool.
	//
	// example:
	//
	// pool-89528023225442****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the basic configuration of the address pool was modified.
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmAddressPoolBasicConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolBasicConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponseBody) SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolBasicConfigResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponseBody) SetRequestId(v string) *UpdateCloudGtmAddressPoolBasicConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponseBody) SetSuccess(v bool) *UpdateCloudGtmAddressPoolBasicConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
