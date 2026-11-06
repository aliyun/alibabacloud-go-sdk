// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribeDataProcessJobTemplatesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeDataProcessJobTemplatesRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeDataProcessJobTemplatesRequest
	GetRegionId() *string
	SetTemplateId(v string) *DescribeDataProcessJobTemplatesRequest
	GetTemplateId() *string
	SetTemplateType(v string) *DescribeDataProcessJobTemplatesRequest
	GetTemplateType() *string
}

type DescribeDataProcessJobTemplatesRequest struct {
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the template resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task template ID used for filtering.
	//
	// example:
	//
	// tpl-123456
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The node templatetype used for filtering. Valid value: image.
	//
	// example:
	//
	// image
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDataProcessJobTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobTemplatesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDataProcessJobTemplatesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDataProcessJobTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataProcessJobTemplatesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeDataProcessJobTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeDataProcessJobTemplatesRequest) SetPageNumber(v string) *DescribeDataProcessJobTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesRequest) SetPageSize(v string) *DescribeDataProcessJobTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesRequest) SetRegionId(v string) *DescribeDataProcessJobTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesRequest) SetTemplateId(v string) *DescribeDataProcessJobTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesRequest) SetTemplateType(v string) *DescribeDataProcessJobTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *DescribeDataProcessJobTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
