// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressId(v string) *CreateCloudGtmAddressResponseBody
	GetAddressId() *string
	SetRequestId(v string) *CreateCloudGtmAddressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCloudGtmAddressResponseBody
	GetSuccess() *bool
}

type CreateCloudGtmAddressResponseBody struct {
	// The address ID. This ID uniquely identifies the address.
	//
	// example:
	//
	// addr-8951821811436**192
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the address creation operation is successful:
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

func (s CreateCloudGtmAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmAddressResponseBody) GetAddressId() *string {
	return s.AddressId
}

func (s *CreateCloudGtmAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudGtmAddressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCloudGtmAddressResponseBody) SetAddressId(v string) *CreateCloudGtmAddressResponseBody {
	s.AddressId = &v
	return s
}

func (s *CreateCloudGtmAddressResponseBody) SetRequestId(v string) *CreateCloudGtmAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudGtmAddressResponseBody) SetSuccess(v bool) *CreateCloudGtmAddressResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCloudGtmAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
