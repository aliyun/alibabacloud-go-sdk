// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GetQualityRuleRequest
	GetProjectId() *int64
	SetProjectName(v string) *GetQualityRuleRequest
	GetProjectName() *string
	SetRuleId(v int64) *GetQualityRuleRequest
	GetRuleId() *int64
}

type GetQualityRuleRequest struct {
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The monitoring rule ID. You can call the [ListQualityRules](https://help.aliyun.com/document_detail/173995.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityRuleRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetQualityRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetQualityRuleRequest) SetProjectId(v int64) *GetQualityRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *GetQualityRuleRequest) SetProjectName(v string) *GetQualityRuleRequest {
	s.ProjectName = &v
	return s
}

func (s *GetQualityRuleRequest) SetRuleId(v int64) *GetQualityRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
