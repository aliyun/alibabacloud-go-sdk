// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeWorkZonesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeWorkZonesRequest
	GetPageSize() *int32
	SetTenantId(v string) *DescribeWorkZonesRequest
	GetTenantId() *string
	SetZoneIdList(v []*string) *DescribeWorkZonesRequest
	GetZoneIdList() []*string
	SetZoneNameList(v []*string) *DescribeWorkZonesRequest
	GetZoneNameList() []*string
}

type DescribeWorkZonesRequest struct {
	PageNumber   *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId     *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneIdList   []*string `json:"ZoneIdList,omitempty" xml:"ZoneIdList,omitempty" type:"Repeated"`
	ZoneNameList []*string `json:"ZoneNameList,omitempty" xml:"ZoneNameList,omitempty" type:"Repeated"`
}

func (s DescribeWorkZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeWorkZonesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWorkZonesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeWorkZonesRequest) GetZoneIdList() []*string {
	return s.ZoneIdList
}

func (s *DescribeWorkZonesRequest) GetZoneNameList() []*string {
	return s.ZoneNameList
}

func (s *DescribeWorkZonesRequest) SetPageNumber(v int32) *DescribeWorkZonesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWorkZonesRequest) SetPageSize(v int32) *DescribeWorkZonesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWorkZonesRequest) SetTenantId(v string) *DescribeWorkZonesRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeWorkZonesRequest) SetZoneIdList(v []*string) *DescribeWorkZonesRequest {
	s.ZoneIdList = v
	return s
}

func (s *DescribeWorkZonesRequest) SetZoneNameList(v []*string) *DescribeWorkZonesRequest {
	s.ZoneNameList = v
	return s
}

func (s *DescribeWorkZonesRequest) Validate() error {
	return dara.Validate(s)
}
