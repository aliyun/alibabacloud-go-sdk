// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *DeleteQualityRuleRequest
	GetProjectId() *int64
	SetProjectName(v string) *DeleteQualityRuleRequest
	GetProjectName() *string
	SetRuleId(v int64) *DeleteQualityRuleRequest
	GetRuleId() *int64
}

type DeleteQualityRuleRequest struct {
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the compute engine or data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The monitoring rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteQualityRuleRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteQualityRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteQualityRuleRequest) SetProjectId(v int64) *DeleteQualityRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteQualityRuleRequest) SetProjectName(v string) *DeleteQualityRuleRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteQualityRuleRequest) SetRuleId(v int64) *DeleteQualityRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
