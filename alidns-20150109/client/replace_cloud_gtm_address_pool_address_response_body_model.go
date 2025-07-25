// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceCloudGtmAddressPoolAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolId(v string) *ReplaceCloudGtmAddressPoolAddressResponseBody
	GetAddressPoolId() *string
	SetRequestId(v string) *ReplaceCloudGtmAddressPoolAddressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReplaceCloudGtmAddressPoolAddressResponseBody
	GetSuccess() *bool
}

type ReplaceCloudGtmAddressPoolAddressResponseBody struct {
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89618921167339**24
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful:
	//
	// - true: Successful. - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReplaceCloudGtmAddressPoolAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmAddressPoolAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmAddressPoolAddressResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *ReplaceCloudGtmAddressPoolAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceCloudGtmAddressPoolAddressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReplaceCloudGtmAddressPoolAddressResponseBody) SetAddressPoolId(v string) *ReplaceCloudGtmAddressPoolAddressResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressResponseBody) SetRequestId(v string) *ReplaceCloudGtmAddressPoolAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressResponseBody) SetSuccess(v bool) *ReplaceCloudGtmAddressPoolAddressResponseBody {
	s.Success = &v
	return s
}

func (s *ReplaceCloudGtmAddressPoolAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
