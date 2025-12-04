// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateElasticNetworkInterfaceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateElasticNetworkInterfaceRequest
	GetDescription() *string
	SetEnableJumboFrame(v bool) *CreateElasticNetworkInterfaceRequest
	GetEnableJumboFrame() *bool
	SetNodeId(v string) *CreateElasticNetworkInterfaceRequest
	GetNodeId() *string
	SetRegionId(v string) *CreateElasticNetworkInterfaceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateElasticNetworkInterfaceRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreateElasticNetworkInterfaceRequest
	GetSecurityGroupId() *string
	SetTag(v []*CreateElasticNetworkInterfaceRequestTag) *CreateElasticNetworkInterfaceRequest
	GetTag() []*CreateElasticNetworkInterfaceRequestTag
	SetVSwitchId(v string) *CreateElasticNetworkInterfaceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateElasticNetworkInterfaceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateElasticNetworkInterfaceRequest
	GetZoneId() *string
}

type CreateElasticNetworkInterfaceRequest struct {
	// The POP API is not ignored by default and is used for idempotence.
	//
	// example:
	//
	// 3fd79d62-ab1d-11ec-9a53-0242ac110004
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the response code.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether to enable the jumbo frame capability
	//
	// example:
	//
	// True
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// The ID of the Lingjun node.
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Resource GroupId.
	//
	// example:
	//
	// rg-acfmxhucx5ewuwy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz9fj2s3o21nw2****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The resource tags.
	Tag []*CreateElasticNetworkInterfaceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-t4nahb0pxck****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// >  If the NodeId parameter is empty, the VpcId parameter is required. If the NodeId parameter is not empty, the VpcId parameter is optional.
	//
	// example:
	//
	// vpc-uf6aa4ddo97fr****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateElasticNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateElasticNetworkInterfaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateElasticNetworkInterfaceRequest) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *CreateElasticNetworkInterfaceRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateElasticNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateElasticNetworkInterfaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateElasticNetworkInterfaceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateElasticNetworkInterfaceRequest) GetTag() []*CreateElasticNetworkInterfaceRequestTag {
	return s.Tag
}

func (s *CreateElasticNetworkInterfaceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateElasticNetworkInterfaceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateElasticNetworkInterfaceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateElasticNetworkInterfaceRequest) SetClientToken(v string) *CreateElasticNetworkInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetDescription(v string) *CreateElasticNetworkInterfaceRequest {
	s.Description = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetEnableJumboFrame(v bool) *CreateElasticNetworkInterfaceRequest {
	s.EnableJumboFrame = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetNodeId(v string) *CreateElasticNetworkInterfaceRequest {
	s.NodeId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetRegionId(v string) *CreateElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetResourceGroupId(v string) *CreateElasticNetworkInterfaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetSecurityGroupId(v string) *CreateElasticNetworkInterfaceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetTag(v []*CreateElasticNetworkInterfaceRequestTag) *CreateElasticNetworkInterfaceRequest {
	s.Tag = v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetVSwitchId(v string) *CreateElasticNetworkInterfaceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetVpcId(v string) *CreateElasticNetworkInterfaceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) SetZoneId(v string) *CreateElasticNetworkInterfaceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequest) Validate() error {
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

type CreateElasticNetworkInterfaceRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateElasticNetworkInterfaceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticNetworkInterfaceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateElasticNetworkInterfaceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateElasticNetworkInterfaceRequestTag) SetKey(v string) *CreateElasticNetworkInterfaceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequestTag) SetValue(v string) *CreateElasticNetworkInterfaceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateElasticNetworkInterfaceRequestTag) Validate() error {
	return dara.Validate(s)
}
