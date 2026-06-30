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
	SetEndpointGroupRegion(v string) *ListEndpointGroupsRequest
	GetEndpointGroupRegion() *string
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
	// The ID of the Global Accelerator instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Whether to enable the access log. Valid values:
	//
	// - **on**: enables the access log.
	//
	// - **off*	- (default): disables the access log.
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
	EndpointGroupId     *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group. Valid values:
	//
	// - **default**: a default endpoint group.
	//
	// - **virtual**: a virtual endpoint group.
	//
	// - If you omit this parameter, the operation returns all default and virtual endpoint groups.
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
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Global Accelerator instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags used to filter endpoint groups. You can specify up to 20 tags.
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

func (s *ListEndpointGroupsRequest) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
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

func (s *ListEndpointGroupsRequest) SetEndpointGroupRegion(v string) *ListEndpointGroupsRequest {
	s.EndpointGroupRegion = &v
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
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters long and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters long and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
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
