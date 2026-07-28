// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteTargetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteRouteTargetGroupRequest
	GetClientToken() *string
	SetForceDelete(v bool) *DeleteRouteTargetGroupRequest
	GetForceDelete() *bool
	SetRegionId(v string) *DeleteRouteTargetGroupRequest
	GetRegionId() *string
	SetRouteTargetGroupId(v string) *DeleteRouteTargetGroupRequest
	GetRouteTargetGroupId() *string
	SetTag(v []*DeleteRouteTargetGroupRequestTag) *DeleteRouteTargetGroupRequest
	GetTag() []*DeleteRouteTargetGroupRequestTag
}

type DeleteRouteTargetGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client generates the value. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may vary for each API request.
	//
	// example:
	//
	// d7d24a21-f4ba-4454-9173-b3****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ForceDelete *bool   `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	// The ID of the region where the resource group resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the routing target group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtg-xxxx
	RouteTargetGroupId *string `json:"RouteTargetGroupId,omitempty" xml:"RouteTargetGroupId,omitempty"`
	// The tags of the resource.
	Tag []*DeleteRouteTargetGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DeleteRouteTargetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteTargetGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteTargetGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRouteTargetGroupRequest) GetForceDelete() *bool {
	return s.ForceDelete
}

func (s *DeleteRouteTargetGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteTargetGroupRequest) GetRouteTargetGroupId() *string {
	return s.RouteTargetGroupId
}

func (s *DeleteRouteTargetGroupRequest) GetTag() []*DeleteRouteTargetGroupRequestTag {
	return s.Tag
}

func (s *DeleteRouteTargetGroupRequest) SetClientToken(v string) *DeleteRouteTargetGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRouteTargetGroupRequest) SetForceDelete(v bool) *DeleteRouteTargetGroupRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteRouteTargetGroupRequest) SetRegionId(v string) *DeleteRouteTargetGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteTargetGroupRequest) SetRouteTargetGroupId(v string) *DeleteRouteTargetGroupRequest {
	s.RouteTargetGroupId = &v
	return s
}

func (s *DeleteRouteTargetGroupRequest) SetTag(v []*DeleteRouteTargetGroupRequestTag) *DeleteRouteTargetGroupRequest {
	s.Tag = v
	return s
}

func (s *DeleteRouteTargetGroupRequest) Validate() error {
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

type DeleteRouteTargetGroupRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteRouteTargetGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteTargetGroupRequestTag) GoString() string {
	return s.String()
}

func (s *DeleteRouteTargetGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *DeleteRouteTargetGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *DeleteRouteTargetGroupRequestTag) SetKey(v string) *DeleteRouteTargetGroupRequestTag {
	s.Key = &v
	return s
}

func (s *DeleteRouteTargetGroupRequestTag) SetValue(v string) *DeleteRouteTargetGroupRequestTag {
	s.Value = &v
	return s
}

func (s *DeleteRouteTargetGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
