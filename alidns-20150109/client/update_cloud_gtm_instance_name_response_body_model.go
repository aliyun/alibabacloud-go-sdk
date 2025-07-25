// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmInstanceNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmInstanceNameResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmInstanceNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
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

func (s UpdateCloudGtmInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmInstanceNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmInstanceNameResponseBody) SetRequestId(v string) *UpdateCloudGtmInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmInstanceNameResponseBody) SetSuccess(v bool) *UpdateCloudGtmInstanceNameResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
