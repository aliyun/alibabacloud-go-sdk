// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressManualAvailableStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmAddressManualAvailableStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmAddressManualAvailableStatusResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmAddressManualAvailableStatusResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmAddressManualAvailableStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressManualAvailableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponseBody) SetRequestId(v string) *UpdateCloudGtmAddressManualAvailableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponseBody) SetSuccess(v bool) *UpdateCloudGtmAddressManualAvailableStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
