// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmAddressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmAddressResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmAddressResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Modify address base configuration operation status:
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

func (s UpdateCloudGtmAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmAddressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmAddressResponseBody) SetRequestId(v string) *UpdateCloudGtmAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmAddressResponseBody) SetSuccess(v bool) *UpdateCloudGtmAddressResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
