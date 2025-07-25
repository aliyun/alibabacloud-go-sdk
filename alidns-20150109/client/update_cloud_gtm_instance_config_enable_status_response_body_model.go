// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigEnableStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmInstanceConfigEnableStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmInstanceConfigEnableStatusResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmInstanceConfigEnableStatusResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful:
	//
	// - **true**: The call succeeded.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigEnableStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponseBody) SetRequestId(v string) *UpdateCloudGtmInstanceConfigEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponseBody) SetSuccess(v bool) *UpdateCloudGtmInstanceConfigEnableStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
