// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigADConnectorTrustRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v string) *ConfigADConnectorTrustRequest
	GetOfficeSiteId() *string
	SetRdsLicenseDomain(v bool) *ConfigADConnectorTrustRequest
	GetRdsLicenseDomain() *bool
	SetRegionId(v string) *ConfigADConnectorTrustRequest
	GetRegionId() *string
	SetTrustKey(v string) *ConfigADConnectorTrustRequest
	GetTrustKey() *string
}

type ConfigADConnectorTrustRequest struct {
	// The ID of the enterprise AD office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-778418****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// Specifies whether to configure a trust password for the Remote Desktop Services (RDS) License Domain of the enterprise AD office network.
	//
	// Valid values:
	//
	// 	- true: configures a trust password for the RDS License Domain of the AD office network.
	//
	// 	- false: configures a trust password for a regular enterprise AD office network.
	//
	// example:
	//
	// true
	RdsLicenseDomain *bool `json:"RdsLicenseDomain,omitempty" xml:"RdsLicenseDomain,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The trust password. You can specify the password when you configure a trust relationship between the AD domain and the ecd.acs domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// password123***
	TrustKey *string `json:"TrustKey,omitempty" xml:"TrustKey,omitempty"`
}

func (s ConfigADConnectorTrustRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigADConnectorTrustRequest) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorTrustRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ConfigADConnectorTrustRequest) GetRdsLicenseDomain() *bool {
	return s.RdsLicenseDomain
}

func (s *ConfigADConnectorTrustRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigADConnectorTrustRequest) GetTrustKey() *string {
	return s.TrustKey
}

func (s *ConfigADConnectorTrustRequest) SetOfficeSiteId(v string) *ConfigADConnectorTrustRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ConfigADConnectorTrustRequest) SetRdsLicenseDomain(v bool) *ConfigADConnectorTrustRequest {
	s.RdsLicenseDomain = &v
	return s
}

func (s *ConfigADConnectorTrustRequest) SetRegionId(v string) *ConfigADConnectorTrustRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigADConnectorTrustRequest) SetTrustKey(v string) *ConfigADConnectorTrustRequest {
	s.TrustKey = &v
	return s
}

func (s *ConfigADConnectorTrustRequest) Validate() error {
	return dara.Validate(s)
}
