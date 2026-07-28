// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRouteTargetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetRouteTargetGroupRequest
	GetClientToken() *string
	SetRegionId(v string) *GetRouteTargetGroupRequest
	GetRegionId() *string
	SetRouteTargetGroupId(v string) *GetRouteTargetGroupRequest
	GetRouteTargetGroupId() *string
	SetTag(v []*GetRouteTargetGroupRequestTag) *GetRouteTargetGroupRequest
	GetTag() []*GetRouteTargetGroupRequestTag
}

type GetRouteTargetGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region to which the route target group belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the routing target group member.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtg-xxxx
	RouteTargetGroupId *string `json:"RouteTargetGroupId,omitempty" xml:"RouteTargetGroupId,omitempty"`
	// The tag information.
	Tag []*GetRouteTargetGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetRouteTargetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRouteTargetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetRouteTargetGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetRouteTargetGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRouteTargetGroupRequest) GetRouteTargetGroupId() *string {
	return s.RouteTargetGroupId
}

func (s *GetRouteTargetGroupRequest) GetTag() []*GetRouteTargetGroupRequestTag {
	return s.Tag
}

func (s *GetRouteTargetGroupRequest) SetClientToken(v string) *GetRouteTargetGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *GetRouteTargetGroupRequest) SetRegionId(v string) *GetRouteTargetGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetRouteTargetGroupRequest) SetRouteTargetGroupId(v string) *GetRouteTargetGroupRequest {
	s.RouteTargetGroupId = &v
	return s
}

func (s *GetRouteTargetGroupRequest) SetTag(v []*GetRouteTargetGroupRequestTag) *GetRouteTargetGroupRequest {
	s.Tag = v
	return s
}

func (s *GetRouteTargetGroupRequest) Validate() error {
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

type GetRouteTargetGroupRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRouteTargetGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetRouteTargetGroupRequestTag) GoString() string {
	return s.String()
}

func (s *GetRouteTargetGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetRouteTargetGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetRouteTargetGroupRequestTag) SetKey(v string) *GetRouteTargetGroupRequestTag {
	s.Key = &v
	return s
}

func (s *GetRouteTargetGroupRequestTag) SetValue(v string) *GetRouteTargetGroupRequestTag {
	s.Value = &v
	return s
}

func (s *GetRouteTargetGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
