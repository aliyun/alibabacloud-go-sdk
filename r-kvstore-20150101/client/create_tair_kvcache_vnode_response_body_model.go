// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTairKVCacheVNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateTairKVCacheVNodeResponseBody
	GetInstanceId() *string
	SetNodeId(v string) *CreateTairKVCacheVNodeResponseBody
	GetNodeId() *string
	SetOrderId(v int64) *CreateTairKVCacheVNodeResponseBody
	GetOrderId() *int64
	SetRegionId(v string) *CreateTairKVCacheVNodeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *CreateTairKVCacheVNodeResponseBody
	GetRequestId() *string
	SetVClusterId(v string) *CreateTairKVCacheVNodeResponseBody
	GetVClusterId() *string
	SetVkName(v string) *CreateTairKVCacheVNodeResponseBody
	GetVkName() *string
	SetZoneId(v string) *CreateTairKVCacheVNodeResponseBody
	GetZoneId() *string
}

type CreateTairKVCacheVNodeResponseBody struct {
	// example:
	//
	// tv-2zeb1ce76fee****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// vn-03a49876edb****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 20905403119****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2BE6E619-A657-42E3-AD2D-18F8428A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// vc-16965a9267*****-*****
	VClusterId *string `json:"VClusterId,omitempty" xml:"VClusterId,omitempty"`
	// example:
	//
	// tk-2ze4bba3c8fe****
	VkName *string `json:"VkName,omitempty" xml:"VkName,omitempty"`
	// example:
	//
	// cn-beijing-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTairKVCacheVNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTairKVCacheVNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTairKVCacheVNodeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTairKVCacheVNodeResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateTairKVCacheVNodeResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateTairKVCacheVNodeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTairKVCacheVNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTairKVCacheVNodeResponseBody) GetVClusterId() *string {
	return s.VClusterId
}

func (s *CreateTairKVCacheVNodeResponseBody) GetVkName() *string {
	return s.VkName
}

func (s *CreateTairKVCacheVNodeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateTairKVCacheVNodeResponseBody) SetInstanceId(v string) *CreateTairKVCacheVNodeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponseBody) SetNodeId(v string) *CreateTairKVCacheVNodeResponseBody {
	s.NodeId = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponseBody) SetOrderId(v int64) *CreateTairKVCacheVNodeResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponseBody) SetRegionId(v string) *CreateTairKVCacheVNodeResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponseBody) SetRequestId(v string) *CreateTairKVCacheVNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponseBody) SetVClusterId(v string) *CreateTairKVCacheVNodeResponseBody {
	s.VClusterId = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponseBody) SetVkName(v string) *CreateTairKVCacheVNodeResponseBody {
	s.VkName = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponseBody) SetZoneId(v string) *CreateTairKVCacheVNodeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateTairKVCacheVNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
