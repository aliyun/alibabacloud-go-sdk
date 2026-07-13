// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v bool) *CreateCloudGtmInstanceConfigResponseBody
	GetConfigId() *bool
	SetInstanceId(v string) *CreateCloudGtmInstanceConfigResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateCloudGtmInstanceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCloudGtmInstanceConfigResponseBody
	GetSuccess() *bool
}

type CreateCloudGtmInstanceConfigResponseBody struct {
	// The ID of the domain name instance configuration. A GTM instance can have both an A record and an AAAA record for the same access domain name. In this case, the GTM instance has two domain name instance configurations. The ConfigId uniquely identifies each configuration.
	//
	// example:
	//
	// config-000**1
	ConfigId *bool `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// gtm-cn-xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful:
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

func (s CreateCloudGtmInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmInstanceConfigResponseBody) GetConfigId() *bool {
	return s.ConfigId
}

func (s *CreateCloudGtmInstanceConfigResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloudGtmInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudGtmInstanceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCloudGtmInstanceConfigResponseBody) SetConfigId(v bool) *CreateCloudGtmInstanceConfigResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigResponseBody) SetInstanceId(v string) *CreateCloudGtmInstanceConfigResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigResponseBody) SetRequestId(v string) *CreateCloudGtmInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigResponseBody) SetSuccess(v bool) *CreateCloudGtmInstanceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
