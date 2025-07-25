// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolId(v string) *CreateCloudGtmAddressPoolResponseBody
	GetAddressPoolId() *string
	SetRequestId(v string) *CreateCloudGtmAddressPoolResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCloudGtmAddressPoolResponseBody
	GetSuccess() *bool
}

type CreateCloudGtmAddressPoolResponseBody struct {
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
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the address pool creation operation was successful:
	//
	// - true: Operation was successful
	//
	// - false: Operation was failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCloudGtmAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmAddressPoolResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *CreateCloudGtmAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudGtmAddressPoolResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCloudGtmAddressPoolResponseBody) SetAddressPoolId(v string) *CreateCloudGtmAddressPoolResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *CreateCloudGtmAddressPoolResponseBody) SetRequestId(v string) *CreateCloudGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudGtmAddressPoolResponseBody) SetSuccess(v bool) *CreateCloudGtmAddressPoolResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCloudGtmAddressPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
