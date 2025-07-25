// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmAddressRemarkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmAddressRemarkResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmAddressRemarkResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the modification operation was successful:
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

func (s UpdateCloudGtmAddressRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmAddressRemarkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmAddressRemarkResponseBody) SetRequestId(v string) *UpdateCloudGtmAddressRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmAddressRemarkResponseBody) SetSuccess(v bool) *UpdateCloudGtmAddressRemarkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmAddressRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
