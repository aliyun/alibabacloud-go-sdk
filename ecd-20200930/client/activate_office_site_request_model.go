// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateOfficeSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v string) *ActivateOfficeSiteRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ActivateOfficeSiteRequest
	GetRegionId() *string
}

type ActivateOfficeSiteRequest struct {
	// The ID of the locked convenience office network.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-803704****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ActivateOfficeSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateOfficeSiteRequest) GoString() string {
	return s.String()
}

func (s *ActivateOfficeSiteRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ActivateOfficeSiteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ActivateOfficeSiteRequest) SetOfficeSiteId(v string) *ActivateOfficeSiteRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ActivateOfficeSiteRequest) SetRegionId(v string) *ActivateOfficeSiteRequest {
	s.RegionId = &v
	return s
}

func (s *ActivateOfficeSiteRequest) Validate() error {
	return dara.Validate(s)
}
