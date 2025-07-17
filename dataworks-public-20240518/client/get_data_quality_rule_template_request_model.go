// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityRuleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataQualityRuleTemplateRequest
	GetCode() *string
}

type GetDataQualityRuleTemplateRequest struct {
	// The code for the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER_DEFINED:123
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDataQualityRuleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleTemplateRequest) GetCode() *string {
	return s.Code
}

func (s *GetDataQualityRuleTemplateRequest) SetCode(v string) *GetDataQualityRuleTemplateRequest {
	s.Code = &v
	return s
}

func (s *GetDataQualityRuleTemplateRequest) Validate() error {
	return dara.Validate(s)
}
