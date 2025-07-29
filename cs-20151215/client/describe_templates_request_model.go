// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int64) *DescribeTemplatesRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeTemplatesRequest
	GetPageSize() *int64
	SetTemplateType(v string) *DescribeTemplatesRequest
	GetTemplateType() *string
}

type DescribeTemplatesRequest struct {
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The type of template. This parameter can be set to a custom value.
	//
	// 	- If the parameter is set to `kubernetes`, the template is displayed on the Templates page in the console.
	//
	// 	- If you set the parameter to `compose`, the template is not displayed on the Templates page in the console.
	//
	// Default value: `kubernetes`.
	//
	// example:
	//
	// kubernetes
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
}

func (s DescribeTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeTemplatesRequest) SetPageNum(v int64) *DescribeTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeTemplatesRequest) SetPageSize(v int64) *DescribeTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatesRequest) SetTemplateType(v string) *DescribeTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
