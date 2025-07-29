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
	// The type of template. The value can be a custom value.
	//
	// 	- If the parameter is set to `kubernetes`, the template is displayed on the Templates page in the console.
	//
	// 	- If the parameter is set to `compose`, the template is displayed on the Container Service - Swarm page in the console. Container Service for Swarm is deprecated.
	//
	// 	- If the value of the parameter is not `kubernetes`, the template is not displayed on the Templates page in the console. We recommend that you set the parameter to `kubernetes`.
	//
	// Default value: `kubernetes`.
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
