// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmInstanceConfigRemarkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmInstanceConfigRemarkResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmInstanceConfigRemarkResponseBody struct {
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponseBody) SetRequestId(v string) *UpdateCloudGtmInstanceConfigRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponseBody) SetSuccess(v bool) *UpdateCloudGtmInstanceConfigRemarkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
