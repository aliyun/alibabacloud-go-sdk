// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetQualityRuleTaskRequest
	GetOpTenantId() *int64
	SetRuleTaskId(v int64) *GetQualityRuleTaskRequest
	GetRuleTaskId() *int64
}

type GetQualityRuleTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	RuleTaskId *int64 `json:"RuleTaskId,omitempty" xml:"RuleTaskId,omitempty"`
}

func (s GetQualityRuleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTaskRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityRuleTaskRequest) GetRuleTaskId() *int64 {
	return s.RuleTaskId
}

func (s *GetQualityRuleTaskRequest) SetOpTenantId(v int64) *GetQualityRuleTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityRuleTaskRequest) SetRuleTaskId(v int64) *GetQualityRuleTaskRequest {
	s.RuleTaskId = &v
	return s
}

func (s *GetQualityRuleTaskRequest) Validate() error {
	return dara.Validate(s)
}
