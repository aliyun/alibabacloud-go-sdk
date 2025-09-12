// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2InstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyLindormV2InstanceResponseBody
	GetAccessDeniedDetail() *string
	SetInstanceId(v string) *ModifyLindormV2InstanceResponseBody
	GetInstanceId() *string
	SetOrderId(v int64) *ModifyLindormV2InstanceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ModifyLindormV2InstanceResponseBody
	GetRequestId() *string
}

type ModifyLindormV2InstanceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId            *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLindormV2InstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyLindormV2InstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyLindormV2InstanceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ModifyLindormV2InstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLindormV2InstanceResponseBody) SetAccessDeniedDetail(v string) *ModifyLindormV2InstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyLindormV2InstanceResponseBody) SetInstanceId(v string) *ModifyLindormV2InstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyLindormV2InstanceResponseBody) SetOrderId(v int64) *ModifyLindormV2InstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyLindormV2InstanceResponseBody) SetRequestId(v string) *ModifyLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLindormV2InstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
