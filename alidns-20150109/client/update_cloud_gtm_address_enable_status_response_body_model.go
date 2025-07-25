// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressEnableStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmAddressEnableStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmAddressEnableStatusResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmAddressEnableStatusResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// 	- true: The operation was successful.
	//
	// 	- false: The operation was failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmAddressEnableStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressEnableStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmAddressEnableStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmAddressEnableStatusResponseBody) SetRequestId(v string) *UpdateCloudGtmAddressEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusResponseBody) SetSuccess(v bool) *UpdateCloudGtmAddressEnableStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
