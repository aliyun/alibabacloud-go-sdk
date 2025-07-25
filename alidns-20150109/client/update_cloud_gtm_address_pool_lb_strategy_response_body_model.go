// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolLbStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolLbStrategyResponseBody
	GetAddressPoolId() *string
	SetRequestId(v string) *UpdateCloudGtmAddressPoolLbStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmAddressPoolLbStrategyResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmAddressPoolLbStrategyResponseBody struct {
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
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
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

func (s UpdateCloudGtmAddressPoolLbStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolLbStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponseBody) SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolLbStrategyResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponseBody) SetRequestId(v string) *UpdateCloudGtmAddressPoolLbStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponseBody) SetSuccess(v bool) *UpdateCloudGtmAddressPoolLbStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolLbStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
