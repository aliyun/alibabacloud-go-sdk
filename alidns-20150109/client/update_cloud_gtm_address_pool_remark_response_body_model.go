// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolRemarkResponseBody
	GetAddressPoolId() *string
	SetRequestId(v string) *UpdateCloudGtmAddressPoolRemarkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmAddressPoolRemarkResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmAddressPoolRemarkResponseBody struct {
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89528023225442**16
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the modification operation was successful:
	//
	// - true: Operation successful
	//
	// - false: Operation failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmAddressPoolRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolRemarkResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *UpdateCloudGtmAddressPoolRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmAddressPoolRemarkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmAddressPoolRemarkResponseBody) SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolRemarkResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkResponseBody) SetRequestId(v string) *UpdateCloudGtmAddressPoolRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkResponseBody) SetSuccess(v bool) *UpdateCloudGtmAddressPoolRemarkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
