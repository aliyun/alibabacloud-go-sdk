// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnLabelCountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelList(v []*string) *DescribeSnLabelCountsRequest
	GetLabelList() []*string
	SetTenantId(v string) *DescribeSnLabelCountsRequest
	GetTenantId() *string
	SetZoneId(v string) *DescribeSnLabelCountsRequest
	GetZoneId() *string
	SetZoneName(v string) *DescribeSnLabelCountsRequest
	GetZoneName() *string
}

type DescribeSnLabelCountsRequest struct {
	LabelList []*string `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	TenantId  *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId    *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName  *string   `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeSnLabelCountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnLabelCountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsRequest) GetLabelList() []*string {
	return s.LabelList
}

func (s *DescribeSnLabelCountsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeSnLabelCountsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSnLabelCountsRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeSnLabelCountsRequest) SetLabelList(v []*string) *DescribeSnLabelCountsRequest {
	s.LabelList = v
	return s
}

func (s *DescribeSnLabelCountsRequest) SetTenantId(v string) *DescribeSnLabelCountsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeSnLabelCountsRequest) SetZoneId(v string) *DescribeSnLabelCountsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeSnLabelCountsRequest) SetZoneName(v string) *DescribeSnLabelCountsRequest {
	s.ZoneName = &v
	return s
}

func (s *DescribeSnLabelCountsRequest) Validate() error {
	return dara.Validate(s)
}
