// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateNetworkServiceRequest
	GetClientToken() *string
	SetName(v string) *CreateNetworkServiceRequest
	GetName() *string
	SetSecurityGroupId(v string) *CreateNetworkServiceRequest
	GetSecurityGroupId() *string
	SetType(v string) *CreateNetworkServiceRequest
	GetType() *string
	SetVpcId(v string) *CreateNetworkServiceRequest
	GetVpcId() *string
	SetVswitchIds(v []*string) *CreateNetworkServiceRequest
	GetVswitchIds() []*string
	SetRegionId(v string) *CreateNetworkServiceRequest
	GetRegionId() *string
}

type CreateNetworkServiceRequest struct {
	// example:
	//
	// acaf8f54-d40e-4c33-a900-f6c1b345cb47
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// securityGroupId
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// example:
	//
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// VPC id。
	//
	// example:
	//
	// vpc-bp1g14f566kbk8jex****
	VpcId      *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	VswitchIds []*string `json:"vswitchIds,omitempty" xml:"vswitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateNetworkServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNetworkServiceRequest) GetName() *string {
	return s.Name
}

func (s *CreateNetworkServiceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateNetworkServiceRequest) GetType() *string {
	return s.Type
}

func (s *CreateNetworkServiceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNetworkServiceRequest) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *CreateNetworkServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkServiceRequest) SetClientToken(v string) *CreateNetworkServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkServiceRequest) SetName(v string) *CreateNetworkServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateNetworkServiceRequest) SetSecurityGroupId(v string) *CreateNetworkServiceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateNetworkServiceRequest) SetType(v string) *CreateNetworkServiceRequest {
	s.Type = &v
	return s
}

func (s *CreateNetworkServiceRequest) SetVpcId(v string) *CreateNetworkServiceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkServiceRequest) SetVswitchIds(v []*string) *CreateNetworkServiceRequest {
	s.VswitchIds = v
	return s
}

func (s *CreateNetworkServiceRequest) SetRegionId(v string) *CreateNetworkServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkServiceRequest) Validate() error {
	return dara.Validate(s)
}
