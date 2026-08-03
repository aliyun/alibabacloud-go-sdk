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
	// The page number. The value starts from 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of results to return.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the template. This operation performs a case-insensitive, fuzzy match. If you do not specify a name, all templates are returned.
	//
	// For example, if you specify `a`, templates named `a1` and `a2` are returned. If you leave this parameter empty, templates named `a1`, `a2`, `b1`, and `c1` are returned.
	//
	// example:
	//
	// example-template
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
