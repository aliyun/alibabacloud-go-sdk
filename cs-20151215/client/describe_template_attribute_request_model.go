// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateType(v string) *DescribeTemplateAttributeRequest
	GetTemplateType() *string
}

type DescribeTemplateAttributeRequest struct {
	// The templatetype.
	//
	// - If you set this parameter to `kubernetes`, the template is displayed on the Orchestration Templates page in the console.
	//
	// - If you leave this parameter empty or set it to other values, the template is not displayed on the Orchestration Templates page in the console.
	//
	// Settings this parameter to `kubernetes` is recommended.
	//
	// example:
	//
	// kubernetes
	TemplateType *string `json:"template_type,omitempty" xml:"template_type,omitempty"`
}

func (s DescribeTemplateAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAttributeRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeTemplateAttributeRequest) SetTemplateType(v string) *DescribeTemplateAttributeRequest {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplateAttributeRequest) Validate() error {
	return dara.Validate(s)
}
