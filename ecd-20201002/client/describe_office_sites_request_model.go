// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfficeSitesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeOfficeSitesRequest
	GetClientId() *string
	SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest
	GetOfficeSiteId() []*string
	SetRegionId(v string) *DescribeOfficeSitesRequest
	GetRegionId() *string
}

type DescribeOfficeSitesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 42f6645a-9c3c-4772-be2a-cc5f5732****
	ClientId     *string   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOfficeSitesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeOfficeSitesRequest) GetOfficeSiteId() []*string {
	return s.OfficeSiteId
}

func (s *DescribeOfficeSitesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOfficeSitesRequest) SetClientId(v string) *DescribeOfficeSitesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

func (s *DescribeOfficeSitesRequest) SetRegionId(v string) *DescribeOfficeSitesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) Validate() error {
	return dara.Validate(s)
}
