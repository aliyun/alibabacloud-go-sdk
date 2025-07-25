// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmInstanceConfigAlertResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmInstanceConfigAlertResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmInstanceConfigAlertResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigAlertResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmInstanceConfigAlertResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmInstanceConfigAlertResponseBody) SetRequestId(v string) *UpdateCloudGtmInstanceConfigAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertResponseBody) SetSuccess(v bool) *UpdateCloudGtmInstanceConfigAlertResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
