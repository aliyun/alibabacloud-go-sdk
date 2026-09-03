// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOfficeSiteSsoStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableSso(v bool) *SetOfficeSiteSsoStatusRequest
	GetEnableSso() *bool
	SetOfficeSiteId(v string) *SetOfficeSiteSsoStatusRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *SetOfficeSiteSsoStatusRequest
	GetRegionId() *string
}

type SetOfficeSiteSsoStatusRequest struct {
	// Specifies whether to enable or shutdown single sign-on (SSO) logon.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableSso *bool `json:"EnableSso,omitempty" xml:"EnableSso,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetOfficeSiteSsoStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetOfficeSiteSsoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetOfficeSiteSsoStatusRequest) GetEnableSso() *bool {
	return s.EnableSso
}

func (s *SetOfficeSiteSsoStatusRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *SetOfficeSiteSsoStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetOfficeSiteSsoStatusRequest) SetEnableSso(v bool) *SetOfficeSiteSsoStatusRequest {
	s.EnableSso = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetOfficeSiteId(v string) *SetOfficeSiteSsoStatusRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) SetRegionId(v string) *SetOfficeSiteSsoStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetOfficeSiteSsoStatusRequest) Validate() error {
	return dara.Validate(s)
}
