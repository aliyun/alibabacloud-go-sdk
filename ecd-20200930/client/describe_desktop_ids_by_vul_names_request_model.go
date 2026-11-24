// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopIdsByVulNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNecessity(v string) *DescribeDesktopIdsByVulNamesRequest
	GetNecessity() *string
	SetOfficeSiteId(v string) *DescribeDesktopIdsByVulNamesRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeDesktopIdsByVulNamesRequest
	GetRegionId() *string
	SetType(v string) *DescribeDesktopIdsByVulNamesRequest
	GetType() *string
	SetVulName(v []*string) *DescribeDesktopIdsByVulNamesRequest
	GetVulName() []*string
}

type DescribeDesktopIdsByVulNamesRequest struct {
	// The necessity level for vulnerability fixing, which indicates the urgency of the vulnerability. Valid values:
	//
	// - asap: high
	//
	// - later: medium
	//
	// - nntf: low
	//
	// example:
	//
	// asap
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vulnerability type. Valid values:
	//
	// - cve: Linux software vulnerability
	//
	// - sys: Windows system vulnerability
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The array of vulnerability names.
	//
	// This parameter is required.
	VulName []*string `json:"VulName,omitempty" xml:"VulName,omitempty" type:"Repeated"`
}

func (s DescribeDesktopIdsByVulNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesRequest) GetNecessity() *string {
	return s.Necessity
}

func (s *DescribeDesktopIdsByVulNamesRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopIdsByVulNamesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopIdsByVulNamesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDesktopIdsByVulNamesRequest) GetVulName() []*string {
	return s.VulName
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetNecessity(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetOfficeSiteId(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetRegionId(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetType(v string) *DescribeDesktopIdsByVulNamesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) SetVulName(v []*string) *DescribeDesktopIdsByVulNamesRequest {
	s.VulName = v
	return s
}

func (s *DescribeDesktopIdsByVulNamesRequest) Validate() error {
	return dara.Validate(s)
}
