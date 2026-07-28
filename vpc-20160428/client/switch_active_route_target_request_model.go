// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchActiveRouteTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SwitchActiveRouteTargetRequest
	GetClientToken() *string
	SetRegionId(v string) *SwitchActiveRouteTargetRequest
	GetRegionId() *string
	SetRouteTargetGroupId(v string) *SwitchActiveRouteTargetRequest
	GetRouteTargetGroupId() *string
	SetTag(v []*SwitchActiveRouteTargetRequestTag) *SwitchActiveRouteTargetRequest
	GetTag() []*SwitchActiveRouteTargetRequestTag
}

type SwitchActiveRouteTargetRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the RequestId value as the ClientToken value. The RequestId value may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the route target group.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the route target group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtg-xxxx
	RouteTargetGroupId *string `json:"RouteTargetGroupId,omitempty" xml:"RouteTargetGroupId,omitempty"`
	// The tags of the resource.
	Tag []*SwitchActiveRouteTargetRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SwitchActiveRouteTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchActiveRouteTargetRequest) GoString() string {
	return s.String()
}

func (s *SwitchActiveRouteTargetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SwitchActiveRouteTargetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SwitchActiveRouteTargetRequest) GetRouteTargetGroupId() *string {
	return s.RouteTargetGroupId
}

func (s *SwitchActiveRouteTargetRequest) GetTag() []*SwitchActiveRouteTargetRequestTag {
	return s.Tag
}

func (s *SwitchActiveRouteTargetRequest) SetClientToken(v string) *SwitchActiveRouteTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *SwitchActiveRouteTargetRequest) SetRegionId(v string) *SwitchActiveRouteTargetRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchActiveRouteTargetRequest) SetRouteTargetGroupId(v string) *SwitchActiveRouteTargetRequest {
	s.RouteTargetGroupId = &v
	return s
}

func (s *SwitchActiveRouteTargetRequest) SetTag(v []*SwitchActiveRouteTargetRequestTag) *SwitchActiveRouteTargetRequest {
	s.Tag = v
	return s
}

func (s *SwitchActiveRouteTargetRequest) Validate() error {
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

type SwitchActiveRouteTargetRequestTag struct {
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

func (s SwitchActiveRouteTargetRequestTag) String() string {
	return dara.Prettify(s)
}

func (s SwitchActiveRouteTargetRequestTag) GoString() string {
	return s.String()
}

func (s *SwitchActiveRouteTargetRequestTag) GetKey() *string {
	return s.Key
}

func (s *SwitchActiveRouteTargetRequestTag) GetValue() *string {
	return s.Value
}

func (s *SwitchActiveRouteTargetRequestTag) SetKey(v string) *SwitchActiveRouteTargetRequestTag {
	s.Key = &v
	return s
}

func (s *SwitchActiveRouteTargetRequestTag) SetValue(v string) *SwitchActiveRouteTargetRequestTag {
	s.Value = &v
	return s
}

func (s *SwitchActiveRouteTargetRequestTag) Validate() error {
	return dara.Validate(s)
}
