// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfficeSitesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v []*string) *DeleteOfficeSitesRequest
	GetOfficeSiteId() []*string
	SetRegionId(v string) *DeleteOfficeSitesRequest
	GetRegionId() *string
}

type DeleteOfficeSitesRequest struct {
	// The IDs of the office networks. You can specify 1 to 100 office networks.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteOfficeSitesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesRequest) GetOfficeSiteId() []*string {
	return s.OfficeSiteId
}

func (s *DeleteOfficeSitesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteOfficeSitesRequest) SetOfficeSiteId(v []*string) *DeleteOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

func (s *DeleteOfficeSitesRequest) SetRegionId(v string) *DeleteOfficeSitesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOfficeSitesRequest) Validate() error {
	return dara.Validate(s)
}
