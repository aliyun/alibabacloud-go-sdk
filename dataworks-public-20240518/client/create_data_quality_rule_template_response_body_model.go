// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityRuleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDataQualityRuleTemplateResponseBody
	GetCode() *string
	SetRequestId(v string) *CreateDataQualityRuleTemplateResponseBody
	GetRequestId() *string
}

type CreateDataQualityRuleTemplateResponseBody struct {
	// The Code of the rule template.
	//
	// example:
	//
	// UserDefined:3001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataQualityRuleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDataQualityRuleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataQualityRuleTemplateResponseBody) SetCode(v string) *CreateDataQualityRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDataQualityRuleTemplateResponseBody) SetRequestId(v string) *CreateDataQualityRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataQualityRuleTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
