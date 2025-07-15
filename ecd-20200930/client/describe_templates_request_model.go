// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *DescribeTemplatesRequest
	GetBizRegionId() *string
	SetBizType(v string) *DescribeTemplatesRequest
	GetBizType() *string
	SetImageId(v string) *DescribeTemplatesRequest
	GetImageId() *string
	SetKeyword(v string) *DescribeTemplatesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DescribeTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTemplatesRequest
	GetPageSize() *int32
	SetProductType(v string) *DescribeTemplatesRequest
	GetProductType() *string
	SetTemplateIds(v []*string) *DescribeTemplatesRequest
	GetTemplateIds() []*string
	SetTemplateName(v string) *DescribeTemplatesRequest
	GetTemplateName() *string
	SetTemplateType(v string) *DescribeTemplatesRequest
	GetTemplateType() *string
}

type DescribeTemplatesRequest struct {
	// example:
	//
	// cn-beijing
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// null
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// m-dnz9xjgbm8*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// CloudDesktop
	ProductType  *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	TemplateIds  []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
	TemplateName *string   `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// USER_TEMPLATE
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeTemplatesRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeTemplatesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeTemplatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTemplatesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeTemplatesRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *DescribeTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeTemplatesRequest) SetBizRegionId(v string) *DescribeTemplatesRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeTemplatesRequest) SetBizType(v string) *DescribeTemplatesRequest {
	s.BizType = &v
	return s
}

func (s *DescribeTemplatesRequest) SetImageId(v string) *DescribeTemplatesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeTemplatesRequest) SetKeyword(v string) *DescribeTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeTemplatesRequest) SetPageNumber(v int32) *DescribeTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTemplatesRequest) SetPageSize(v int32) *DescribeTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesRequest) SetProductType(v string) *DescribeTemplatesRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeTemplatesRequest) SetTemplateIds(v []*string) *DescribeTemplatesRequest {
	s.TemplateIds = v
	return s
}

func (s *DescribeTemplatesRequest) SetTemplateName(v string) *DescribeTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeTemplatesRequest) SetTemplateType(v string) *DescribeTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
