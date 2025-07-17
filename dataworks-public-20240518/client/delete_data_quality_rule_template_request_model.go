// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityRuleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDataQualityRuleTemplateRequest
	GetCode() *string
	SetProjectId(v int64) *DeleteDataQualityRuleTemplateRequest
	GetProjectId() *int64
}

type DeleteDataQualityRuleTemplateRequest struct {
	// The code for the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER_DEFINED:123
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDataQualityRuleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityRuleTemplateRequest) GetCode() *string {
	return s.Code
}

func (s *DeleteDataQualityRuleTemplateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDataQualityRuleTemplateRequest) SetCode(v string) *DeleteDataQualityRuleTemplateRequest {
	s.Code = &v
	return s
}

func (s *DeleteDataQualityRuleTemplateRequest) SetProjectId(v int64) *DeleteDataQualityRuleTemplateRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDataQualityRuleTemplateRequest) Validate() error {
	return dara.Validate(s)
}
