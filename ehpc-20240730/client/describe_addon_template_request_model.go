// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *DescribeAddonTemplateRequest
	GetAddonName() *string
	SetAddonVersion(v string) *DescribeAddonTemplateRequest
	GetAddonVersion() *string
	SetPageNumber(v int64) *DescribeAddonTemplateRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAddonTemplateRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeAddonTemplateRequest
	GetRegionId() *string
	SetZoneId(v string) *DescribeAddonTemplateRequest
	GetZoneId() *string
}

type DescribeAddonTemplateRequest struct {
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// The page number of the page returned. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAddonTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *DescribeAddonTemplateRequest) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *DescribeAddonTemplateRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAddonTemplateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAddonTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAddonTemplateRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAddonTemplateRequest) SetAddonName(v string) *DescribeAddonTemplateRequest {
	s.AddonName = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetAddonVersion(v string) *DescribeAddonTemplateRequest {
	s.AddonVersion = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetPageNumber(v int64) *DescribeAddonTemplateRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetPageSize(v int64) *DescribeAddonTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetRegionId(v string) *DescribeAddonTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAddonTemplateRequest) SetZoneId(v string) *DescribeAddonTemplateRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAddonTemplateRequest) Validate() error {
	return dara.Validate(s)
}
