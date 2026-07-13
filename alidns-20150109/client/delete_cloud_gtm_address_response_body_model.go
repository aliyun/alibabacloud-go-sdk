// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudGtmAddressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCloudGtmAddressResponseBody
	GetSuccess() *bool
}

type DeleteCloudGtmAddressResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
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

func (s DeleteCloudGtmAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudGtmAddressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCloudGtmAddressResponseBody) SetRequestId(v string) *DeleteCloudGtmAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudGtmAddressResponseBody) SetSuccess(v bool) *DeleteCloudGtmAddressResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCloudGtmAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
