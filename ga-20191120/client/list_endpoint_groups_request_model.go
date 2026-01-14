// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListEndpointGroupsRequest
	GetAcceleratorId() *string
	SetAccessLogSwitch(v string) *ListEndpointGroupsRequest
	GetAccessLogSwitch() *string
	SetEndpointGroupId(v string) *ListEndpointGroupsRequest
	GetEndpointGroupId() *string
	SetEndpointGroupType(v string) *ListEndpointGroupsRequest
	GetEndpointGroupType() *string
	SetListenerId(v string) *ListEndpointGroupsRequest
	GetListenerId() *string
	SetPageNumber(v int32) *ListEndpointGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEndpointGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListEndpointGroupsRequest
	GetRegionId() *string
	SetTag(v []*ListEndpointGroupsRequestTag) *ListEndpointGroupsRequest
	GetTag() []*ListEndpointGroupsRequestTag
}

type ListEndpointGroupsRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Specifies whether the access logging feature is enabled. Default value: off. Valid values:
	//
	// 	- **on**: The access logging feature is enabled.
	//
	// 	- **off**: The access logging feature is disabled.
	//
	// example:
	//
	// on
	AccessLogSwitch *string `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp16jdc00bhe97sr5****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The type of the endpoint group. Valid values: Valid values:
	//
	// 	- **default**
	//
	// 	- **virtual**
	//
	// 	- If you leave this parameter empty, all default and virtual endpoint groups are queried.
	//
	// example:
	//
	// virtual
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	// The ID of the listener.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag of the endpoint group.
	//
	// if can be null:
	// false
	Tag []*ListEndpointGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListEndpointGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListEndpointGroupsRequest) GetAccessLogSwitch() *string {
	return s.AccessLogSwitch
}

func (s *ListEndpointGroupsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListEndpointGroupsRequest) GetEndpointGroupType() *string {
	return s.EndpointGroupType
}

func (s *ListEndpointGroupsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListEndpointGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEndpointGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEndpointGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEndpointGroupsRequest) GetTag() []*ListEndpointGroupsRequestTag {
	return s.Tag
}

func (s *ListEndpointGroupsRequest) SetAcceleratorId(v string) *ListEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetAccessLogSwitch(v string) *ListEndpointGroupsRequest {
	s.AccessLogSwitch = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetEndpointGroupId(v string) *ListEndpointGroupsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetEndpointGroupType(v string) *ListEndpointGroupsRequest {
	s.EndpointGroupType = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetListenerId(v string) *ListEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetPageNumber(v int32) *ListEndpointGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetPageSize(v int32) *ListEndpointGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetRegionId(v string) *ListEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEndpointGroupsRequest) SetTag(v []*ListEndpointGroupsRequestTag) *ListEndpointGroupsRequest {
	s.Tag = v
	return s
}

func (s *ListEndpointGroupsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEndpointGroupsRequestTag struct {
	// The tag key of the endpoint group. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the endpoint group. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag values.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEndpointGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListEndpointGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListEndpointGroupsRequestTag) SetKey(v string) *ListEndpointGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListEndpointGroupsRequestTag) SetValue(v string) *ListEndpointGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *ListEndpointGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
