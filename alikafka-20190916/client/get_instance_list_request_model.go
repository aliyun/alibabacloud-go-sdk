// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v []*string) *GetInstanceListRequest
	GetInstanceId() []*string
	SetOrderId(v string) *GetInstanceListRequest
	GetOrderId() *string
	SetRegionId(v string) *GetInstanceListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *GetInstanceListRequest
	GetResourceGroupId() *string
	SetSeries(v string) *GetInstanceListRequest
	GetSeries() *string
	SetTag(v []*GetInstanceListRequestTag) *GetInstanceListRequest
	GetTag() []*GetInstanceListRequestTag
}

type GetInstanceListRequest struct {
	// The IDs of instances.
	//
	// example:
	//
	// alikafka_post-cn-mp91gnw0p***
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The ID of the order. You can obtain the order ID on the [Orders](https://usercenter2-intl.aliyun.com/order/list?pageIndex=1\\&pageSize=20\\&spm=5176.12818093.top-nav.ditem-ord.36f016d0OQFmJa) page in Alibaba Cloud User Center.
	//
	// example:
	//
	// 6072673****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can obtain this ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance version. You can use instance versions to filter different versions of instances. Valid values:
	//
	// 	- v2
	//
	// 	- v3
	//
	// 	- confluent
	//
	// example:
	//
	// v3
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// The tags.
	Tag []*GetInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceListRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *GetInstanceListRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *GetInstanceListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceListRequest) GetSeries() *string {
	return s.Series
}

func (s *GetInstanceListRequest) GetTag() []*GetInstanceListRequestTag {
	return s.Tag
}

func (s *GetInstanceListRequest) SetInstanceId(v []*string) *GetInstanceListRequest {
	s.InstanceId = v
	return s
}

func (s *GetInstanceListRequest) SetOrderId(v string) *GetInstanceListRequest {
	s.OrderId = &v
	return s
}

func (s *GetInstanceListRequest) SetRegionId(v string) *GetInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceListRequest) SetResourceGroupId(v string) *GetInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceListRequest) SetSeries(v string) *GetInstanceListRequest {
	s.Series = &v
	return s
}

func (s *GetInstanceListRequest) SetTag(v []*GetInstanceListRequestTag) *GetInstanceListRequest {
	s.Tag = v
	return s
}

func (s *GetInstanceListRequest) Validate() error {
	return dara.Validate(s)
}

type GetInstanceListRequestTag struct {
	// The tag key.
	//
	// 	- If you leave this parameter empty, the keys of all tags are matched.
	//
	// 	- The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// 	- If you leave Key empty, you must also leave this parameter empty. If you leave this parameter empty, the values of all tags are matched.
	//
	// 	- The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *GetInstanceListRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetInstanceListRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetInstanceListRequestTag) SetKey(v string) *GetInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *GetInstanceListRequestTag) SetValue(v string) *GetInstanceListRequestTag {
	s.Value = &v
	return s
}

func (s *GetInstanceListRequestTag) Validate() error {
	return dara.Validate(s)
}
