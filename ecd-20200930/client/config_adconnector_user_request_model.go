// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigADConnectorUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainPassword(v string) *ConfigADConnectorUserRequest
	GetDomainPassword() *string
	SetDomainUserName(v string) *ConfigADConnectorUserRequest
	GetDomainUserName() *string
	SetOUName(v string) *ConfigADConnectorUserRequest
	GetOUName() *string
	SetOfficeSiteId(v string) *ConfigADConnectorUserRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ConfigADConnectorUserRequest
	GetRegionId() *string
}

type ConfigADConnectorUserRequest struct {
	// The password of the AD user that has the permissions to join computers to domains.
	//
	// This parameter is required.
	//
	// example:
	//
	// testPassword
	DomainPassword *string `json:"DomainPassword,omitempty" xml:"DomainPassword,omitempty"`
	// The username of the AD user that has the permissions to join computers to domains.
	//
	// After the username is configured, the cloud desktops in the same AD workspace are joined to the specified OU.
	//
	// This parameter is required.
	//
	// example:
	//
	// Administrator
	DomainUserName *string `json:"DomainUserName,omitempty" xml:"DomainUserName,omitempty"`
	// The name of the OU in the AD domain. You can call the [ListUserAdOrganizationUnits](https://help.aliyun.com/document_detail/311259.html) to obtain the OU name.
	//
	// example:
	//
	// example.com/Domain Controllers
	OUName *string `json:"OUName,omitempty" xml:"OUName,omitempty"`
	// The ID of the AD workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-778418****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigADConnectorUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigADConnectorUserRequest) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorUserRequest) GetDomainPassword() *string {
	return s.DomainPassword
}

func (s *ConfigADConnectorUserRequest) GetDomainUserName() *string {
	return s.DomainUserName
}

func (s *ConfigADConnectorUserRequest) GetOUName() *string {
	return s.OUName
}

func (s *ConfigADConnectorUserRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ConfigADConnectorUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigADConnectorUserRequest) SetDomainPassword(v string) *ConfigADConnectorUserRequest {
	s.DomainPassword = &v
	return s
}

func (s *ConfigADConnectorUserRequest) SetDomainUserName(v string) *ConfigADConnectorUserRequest {
	s.DomainUserName = &v
	return s
}

func (s *ConfigADConnectorUserRequest) SetOUName(v string) *ConfigADConnectorUserRequest {
	s.OUName = &v
	return s
}

func (s *ConfigADConnectorUserRequest) SetOfficeSiteId(v string) *ConfigADConnectorUserRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ConfigADConnectorUserRequest) SetRegionId(v string) *ConfigADConnectorUserRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigADConnectorUserRequest) Validate() error {
	return dara.Validate(s)
}
