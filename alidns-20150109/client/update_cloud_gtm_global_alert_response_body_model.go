// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmGlobalAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmGlobalAlertResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmGlobalAlertResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmGlobalAlertResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmGlobalAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmGlobalAlertResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmGlobalAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmGlobalAlertResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmGlobalAlertResponseBody) SetRequestId(v string) *UpdateCloudGtmGlobalAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertResponseBody) SetSuccess(v bool) *UpdateCloudGtmGlobalAlertResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertResponseBody) Validate() error {
	return dara.Validate(s)
}
