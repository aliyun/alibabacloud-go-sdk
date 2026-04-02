// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnhancedVpnGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPropagate(v bool) *UpdateEnhancedVpnGatewayResponseBody
	GetAutoPropagate() *bool
	SetCreateTime(v int64) *UpdateEnhancedVpnGatewayResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetDescription() *string
	SetDisasterRecoveryVSwitchId(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetDisasterRecoveryVSwitchId() *string
	SetEnableBgp(v bool) *UpdateEnhancedVpnGatewayResponseBody
	GetEnableBgp() *bool
	SetEniInstanceIds(v *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds) *UpdateEnhancedVpnGatewayResponseBody
	GetEniInstanceIds() *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds
	SetGatewayType(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetGatewayType() *string
	SetName(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetName() *string
	SetNetworkType(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetNetworkType() *string
	SetRequestId(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetStatus() *string
	SetTag(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetTag() *string
	SetTags(v *UpdateEnhancedVpnGatewayResponseBodyTags) *UpdateEnhancedVpnGatewayResponseBody
	GetTags() *UpdateEnhancedVpnGatewayResponseBodyTags
	SetVSwitchId(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetVpcId() *string
	SetVpnGatewayId(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetVpnGatewayId() *string
	SetVpnType(v string) *UpdateEnhancedVpnGatewayResponseBody
	GetVpnType() *string
}

type UpdateEnhancedVpnGatewayResponseBody struct {
	// example:
	//
	// true
	AutoPropagate *bool `json:"AutoPropagate,omitempty" xml:"AutoPropagate,omitempty"`
	// example:
	//
	// 1492753580000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// vsw-p0w95ql6tmr2ludkt****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// example:
	//
	// true
	EnableBgp      *bool                                               `json:"EnableBgp,omitempty" xml:"EnableBgp,omitempty"`
	EniInstanceIds *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds `json:"EniInstanceIds,omitempty" xml:"EniInstanceIds,omitempty" type:"Struct"`
	// example:
	//
	// Enhanced.SiteToSite
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B457
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {"VpnEnableBgp": true}
	Tag  *string                                   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Tags *UpdateEnhancedVpnGatewayResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// vsw-bp1y9ovl1cu9ou4tv****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp1ub1yt9cvakoel****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
	// example:
	//
	// Normal
	VpnType *string `json:"VpnType,omitempty" xml:"VpnType,omitempty"`
}

func (s UpdateEnhancedVpnGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnhancedVpnGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetAutoPropagate() *bool {
	return s.AutoPropagate
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetDisasterRecoveryVSwitchId() *string {
	return s.DisasterRecoveryVSwitchId
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetEnableBgp() *bool {
	return s.EnableBgp
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetEniInstanceIds() *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds {
	return s.EniInstanceIds
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetGatewayType() *string {
	return s.GatewayType
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetTag() *string {
	return s.Tag
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetTags() *UpdateEnhancedVpnGatewayResponseBodyTags {
	return s.Tags
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *UpdateEnhancedVpnGatewayResponseBody) GetVpnType() *string {
	return s.VpnType
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetAutoPropagate(v bool) *UpdateEnhancedVpnGatewayResponseBody {
	s.AutoPropagate = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetCreateTime(v int64) *UpdateEnhancedVpnGatewayResponseBody {
	s.CreateTime = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetDescription(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetDisasterRecoveryVSwitchId(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.DisasterRecoveryVSwitchId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetEnableBgp(v bool) *UpdateEnhancedVpnGatewayResponseBody {
	s.EnableBgp = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetEniInstanceIds(v *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds) *UpdateEnhancedVpnGatewayResponseBody {
	s.EniInstanceIds = v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetGatewayType(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.GatewayType = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetName(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetNetworkType(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.NetworkType = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetRequestId(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetResourceGroupId(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetStatus(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetTag(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.Tag = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetTags(v *UpdateEnhancedVpnGatewayResponseBodyTags) *UpdateEnhancedVpnGatewayResponseBody {
	s.Tags = v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetVSwitchId(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetVpcId(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.VpcId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetVpnGatewayId(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.VpnGatewayId = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) SetVpnType(v string) *UpdateEnhancedVpnGatewayResponseBody {
	s.VpnType = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBody) Validate() error {
	if s.EniInstanceIds != nil {
		if err := s.EniInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds struct {
	EniInstanceId []*string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty" type:"Repeated"`
}

func (s UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds) GoString() string {
	return s.String()
}

func (s *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds) GetEniInstanceId() []*string {
	return s.EniInstanceId
}

func (s *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds) SetEniInstanceId(v []*string) *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds {
	s.EniInstanceId = v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBodyEniInstanceIds) Validate() error {
	return dara.Validate(s)
}

type UpdateEnhancedVpnGatewayResponseBodyTags struct {
	Tag []*UpdateEnhancedVpnGatewayResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s UpdateEnhancedVpnGatewayResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnhancedVpnGatewayResponseBodyTags) GoString() string {
	return s.String()
}

func (s *UpdateEnhancedVpnGatewayResponseBodyTags) GetTag() []*UpdateEnhancedVpnGatewayResponseBodyTagsTag {
	return s.Tag
}

func (s *UpdateEnhancedVpnGatewayResponseBodyTags) SetTag(v []*UpdateEnhancedVpnGatewayResponseBodyTagsTag) *UpdateEnhancedVpnGatewayResponseBodyTags {
	s.Tag = v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBodyTags) Validate() error {
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

type UpdateEnhancedVpnGatewayResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateEnhancedVpnGatewayResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnhancedVpnGatewayResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *UpdateEnhancedVpnGatewayResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *UpdateEnhancedVpnGatewayResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *UpdateEnhancedVpnGatewayResponseBodyTagsTag) SetKey(v string) *UpdateEnhancedVpnGatewayResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBodyTagsTag) SetValue(v string) *UpdateEnhancedVpnGatewayResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *UpdateEnhancedVpnGatewayResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
