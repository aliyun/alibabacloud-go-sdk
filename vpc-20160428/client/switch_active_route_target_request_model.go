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
	// Client token, used to ensure the idempotence of the request. Generate a unique value for this parameter from your client to ensure that it is unique across different requests. The ClientToken only supports ASCII characters. Note: If you do not specify this, the system will automatically use the RequestId of the API request as the ClientToken identifier. The RequestId may differ for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe6****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID where the RouteTargetGroup is located.
	//
	// You can obtain the region ID by calling the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the RouteTargetGroup.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtg-xxxx
	RouteTargetGroupId *string `json:"RouteTargetGroupId,omitempty" xml:"RouteTargetGroupId,omitempty"`
	// Resource tags.
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
	// The key of the resource tag. Up to 20 tag keys are supported. If you need to pass this value, you cannot input an empty string.
	//
	// Each tag key can have up to 128 characters and cannot start with `aliyun` or `acs:`. It also cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the resource tag. Up to 20 tag values are supported. If you need to pass this value, you can input an empty string. A maximum of 128 characters is allowed. The value cannot start with `aliyun` or `acs:`, and it must not contain `http://` or `https://`.
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
