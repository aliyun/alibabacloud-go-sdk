// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolEnableStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolEnableStatusResponseBody
	GetAddressPoolId() *string
	SetRequestId(v string) *UpdateCloudGtmAddressPoolEnableStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmAddressPoolEnableStatusResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmAddressPoolEnableStatusResponseBody struct {
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
	// 89184F33-48A1-4401-9C0F-40E45DB091AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the modification operation was successful:
	//
	// - true: Operation was successful
	//
	// - false: Operation failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmAddressPoolEnableStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponseBody) SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolEnableStatusResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponseBody) SetRequestId(v string) *UpdateCloudGtmAddressPoolEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponseBody) SetSuccess(v bool) *UpdateCloudGtmAddressPoolEnableStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
