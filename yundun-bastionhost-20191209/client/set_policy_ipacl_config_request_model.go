// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyIPAclConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIPAclConfig(v *SetPolicyIPAclConfigRequestIPAclConfig) *SetPolicyIPAclConfigRequest
	GetIPAclConfig() *SetPolicyIPAclConfigRequestIPAclConfig
	SetInstanceId(v string) *SetPolicyIPAclConfigRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyIPAclConfigRequest
	GetPolicyId() *string
	SetProjectId(v int64) *SetPolicyIPAclConfigRequest
	GetProjectId() *int64
	SetRegionId(v string) *SetPolicyIPAclConfigRequest
	GetRegionId() *string
}

type SetPolicyIPAclConfigRequest struct {
	// The source IP address restriction configuration.
	//
	// This parameter is required.
	IPAclConfig *SetPolicyIPAclConfigRequestIPAclConfig `json:"IPAclConfig,omitempty" xml:"IPAclConfig,omitempty" type:"Struct"`
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the bastion host instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyIPAclConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyIPAclConfigRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyIPAclConfigRequest) GetIPAclConfig() *SetPolicyIPAclConfigRequestIPAclConfig {
	return s.IPAclConfig
}

func (s *SetPolicyIPAclConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyIPAclConfigRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyIPAclConfigRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SetPolicyIPAclConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyIPAclConfigRequest) SetIPAclConfig(v *SetPolicyIPAclConfigRequestIPAclConfig) *SetPolicyIPAclConfigRequest {
	s.IPAclConfig = v
	return s
}

func (s *SetPolicyIPAclConfigRequest) SetInstanceId(v string) *SetPolicyIPAclConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyIPAclConfigRequest) SetPolicyId(v string) *SetPolicyIPAclConfigRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyIPAclConfigRequest) SetProjectId(v int64) *SetPolicyIPAclConfigRequest {
	s.ProjectId = &v
	return s
}

func (s *SetPolicyIPAclConfigRequest) SetRegionId(v string) *SetPolicyIPAclConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyIPAclConfigRequest) Validate() error {
	if s.IPAclConfig != nil {
		if err := s.IPAclConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetPolicyIPAclConfigRequestIPAclConfig struct {
	// The source IP address restriction mode. Valid values:
	//
	// - **black**: blacklist mode
	//
	// - **white**: whitelist mode
	//
	// This parameter is required.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The list of restricted source IP addresses.
	//
	// > - When AclType is set to white, the IPs list is required.
	//
	// > - When AclType is set to black but you do not need to set a source IP blacklist, pass an empty list for IPs.
	//
	// This parameter is required.
	IPs []*string `json:"IPs,omitempty" xml:"IPs,omitempty" type:"Repeated"`
}

func (s SetPolicyIPAclConfigRequestIPAclConfig) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyIPAclConfigRequestIPAclConfig) GoString() string {
	return s.String()
}

func (s *SetPolicyIPAclConfigRequestIPAclConfig) GetAclType() *string {
	return s.AclType
}

func (s *SetPolicyIPAclConfigRequestIPAclConfig) GetIPs() []*string {
	return s.IPs
}

func (s *SetPolicyIPAclConfigRequestIPAclConfig) SetAclType(v string) *SetPolicyIPAclConfigRequestIPAclConfig {
	s.AclType = &v
	return s
}

func (s *SetPolicyIPAclConfigRequestIPAclConfig) SetIPs(v []*string) *SetPolicyIPAclConfigRequestIPAclConfig {
	s.IPs = v
	return s
}

func (s *SetPolicyIPAclConfigRequestIPAclConfig) Validate() error {
	return dara.Validate(s)
}
