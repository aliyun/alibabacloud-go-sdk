// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTaskLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetQualityRuleTaskLogRequest
	GetOpTenantId() *int64
	SetRuleTaskId(v int64) *GetQualityRuleTaskLogRequest
	GetRuleTaskId() *int64
}

type GetQualityRuleTaskLogRequest struct {
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

func (s GetQualityRuleTaskLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTaskLogRequest) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTaskLogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetQualityRuleTaskLogRequest) GetRuleTaskId() *int64 {
	return s.RuleTaskId
}

func (s *GetQualityRuleTaskLogRequest) SetOpTenantId(v int64) *GetQualityRuleTaskLogRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetQualityRuleTaskLogRequest) SetRuleTaskId(v int64) *GetQualityRuleTaskLogRequest {
	s.RuleTaskId = &v
	return s
}

func (s *GetQualityRuleTaskLogRequest) Validate() error {
	return dara.Validate(s)
}
