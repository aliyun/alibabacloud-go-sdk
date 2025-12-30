// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeZonesRequest
	GetKeyword() *string
	SetLang(v string) *DescribeZonesRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeZonesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeZonesRequest
	GetPageSize() *int32
	SetQueryRegionId(v string) *DescribeZonesRequest
	GetQueryRegionId() *string
	SetQueryVpcId(v string) *DescribeZonesRequest
	GetQueryVpcId() *string
	SetResourceGroupId(v string) *DescribeZonesRequest
	GetResourceGroupId() *string
	SetResourceTag(v []*DescribeZonesRequestResourceTag) *DescribeZonesRequest
	GetResourceTag() []*DescribeZonesRequestResourceTag
	SetSearchMode(v string) *DescribeZonesRequest
	GetSearchMode() *string
	SetZoneTag(v []*string) *DescribeZonesRequest
	GetZoneTag() []*string
	SetZoneType(v string) *DescribeZonesRequest
	GetZoneType() *string
}

type DescribeZonesRequest struct {
	// The keyword of the zone name. The value is not case-sensitive. You can set SearchMode to LIKE or EXACT. The default value of SearchMode is LIKE.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the virtual private cloud (VPC) associated with the zone.
	//
	// example:
	//
	// cn-hangzhou
	QueryRegionId *string `json:"QueryRegionId,omitempty" xml:"QueryRegionId,omitempty"`
	// The ID of the VPC associated with the zone.
	//
	// example:
	//
	// vpc-f8zvrvr1payllgz38****
	QueryVpcId *string `json:"QueryVpcId,omitempty" xml:"QueryVpcId,omitempty"`
	// The ID of the resource group to which the zone belongs.
	//
	// example:
	//
	// rg-aekz2qj7awz****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the zone.
	ResourceTag []*DescribeZonesRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The search mode. The value of Keyword is the search scope. Valid values:
	//
	// 	- **LIKE*	- (default): fuzzy search
	//
	// 	- **EXACT**: exact search
	//
	// Default value: **LIKE**.
	//
	// example:
	//
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The types of cloud services.
	//
	// example:
	//
	// BLINK
	ZoneTag []*string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty" type:"Repeated"`
	// The zone type. Valid values:
	//
	// 	- **AUTH_ZONE**: authoritative zone
	//
	// 	- **CLOUD_PRODUCT_ZONE**: authoritative zone for cloud services
	//
	// Default value: **AUTH_ZONE**.
	//
	// example:
	//
	// CLOUD_PRODUCT_ZONE
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeZonesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeZonesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeZonesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeZonesRequest) GetQueryRegionId() *string {
	return s.QueryRegionId
}

func (s *DescribeZonesRequest) GetQueryVpcId() *string {
	return s.QueryVpcId
}

func (s *DescribeZonesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeZonesRequest) GetResourceTag() []*DescribeZonesRequestResourceTag {
	return s.ResourceTag
}

func (s *DescribeZonesRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeZonesRequest) GetZoneTag() []*string {
	return s.ZoneTag
}

func (s *DescribeZonesRequest) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeZonesRequest) SetKeyword(v string) *DescribeZonesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeZonesRequest) SetLang(v string) *DescribeZonesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZonesRequest) SetPageNumber(v int32) *DescribeZonesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeZonesRequest) SetPageSize(v int32) *DescribeZonesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeZonesRequest) SetQueryRegionId(v string) *DescribeZonesRequest {
	s.QueryRegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetQueryVpcId(v string) *DescribeZonesRequest {
	s.QueryVpcId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceGroupId(v string) *DescribeZonesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceTag(v []*DescribeZonesRequestResourceTag) *DescribeZonesRequest {
	s.ResourceTag = v
	return s
}

func (s *DescribeZonesRequest) SetSearchMode(v string) *DescribeZonesRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeZonesRequest) SetZoneTag(v []*string) *DescribeZonesRequest {
	s.ZoneTag = v
	return s
}

func (s *DescribeZonesRequest) SetZoneType(v string) *DescribeZonesRequest {
	s.ZoneType = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	if s.ResourceTag != nil {
		for _, item := range s.ResourceTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeZonesRequestResourceTag struct {
	// The key of tag N added to the zone.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the zone.
	//
	// example:
	//
	// daily
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeZonesRequestResourceTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequestResourceTag) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequestResourceTag) GetKey() *string {
	return s.Key
}

func (s *DescribeZonesRequestResourceTag) GetValue() *string {
	return s.Value
}

func (s *DescribeZonesRequestResourceTag) SetKey(v string) *DescribeZonesRequestResourceTag {
	s.Key = &v
	return s
}

func (s *DescribeZonesRequestResourceTag) SetValue(v string) *DescribeZonesRequestResourceTag {
	s.Value = &v
	return s
}

func (s *DescribeZonesRequestResourceTag) Validate() error {
	return dara.Validate(s)
}
