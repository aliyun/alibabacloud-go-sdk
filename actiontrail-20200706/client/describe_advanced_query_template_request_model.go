// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedQueryTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeAdvancedQueryTemplateRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAdvancedQueryTemplateRequest
	GetPageSize() *int64
	SetTemplateName(v string) *DescribeAdvancedQueryTemplateRequest
	GetTemplateName() *string
}

type DescribeAdvancedQueryTemplateRequest struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeAdvancedQueryTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedQueryTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedQueryTemplateRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAdvancedQueryTemplateRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAdvancedQueryTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeAdvancedQueryTemplateRequest) SetPageNumber(v int64) *DescribeAdvancedQueryTemplateRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateRequest) SetPageSize(v int64) *DescribeAdvancedQueryTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateRequest) SetTemplateName(v string) *DescribeAdvancedQueryTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateRequest) Validate() error {
	return dara.Validate(s)
}
