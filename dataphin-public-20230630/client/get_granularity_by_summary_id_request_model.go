// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGranularityBySummaryIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetGranularityBySummaryIdRequest
	GetEnv() *string
	SetModelId(v int64) *GetGranularityBySummaryIdRequest
	GetModelId() *int64
	SetOpTenantId(v int64) *GetGranularityBySummaryIdRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetGranularityBySummaryIdRequest
	GetOpUserId() *string
	SetProjectId(v int64) *GetGranularityBySummaryIdRequest
	GetProjectId() *int64
	SetQueryDraft(v bool) *GetGranularityBySummaryIdRequest
	GetQueryDraft() *bool
}

type GetGranularityBySummaryIdRequest struct {
	// Env
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// ModelId
	//
	// This parameter is required.
	//
	// example:
	//
	// 1978
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// ProjectId
	//
	// This parameter is required.
	//
	// example:
	//
	// 7081229106458752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// QueryDraft
	QueryDraft *bool `json:"QueryDraft,omitempty" xml:"QueryDraft,omitempty"`
}

func (s GetGranularityBySummaryIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGranularityBySummaryIdRequest) GoString() string {
	return s.String()
}

func (s *GetGranularityBySummaryIdRequest) GetEnv() *string {
	return s.Env
}

func (s *GetGranularityBySummaryIdRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *GetGranularityBySummaryIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetGranularityBySummaryIdRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetGranularityBySummaryIdRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetGranularityBySummaryIdRequest) GetQueryDraft() *bool {
	return s.QueryDraft
}

func (s *GetGranularityBySummaryIdRequest) SetEnv(v string) *GetGranularityBySummaryIdRequest {
	s.Env = &v
	return s
}

func (s *GetGranularityBySummaryIdRequest) SetModelId(v int64) *GetGranularityBySummaryIdRequest {
	s.ModelId = &v
	return s
}

func (s *GetGranularityBySummaryIdRequest) SetOpTenantId(v int64) *GetGranularityBySummaryIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetGranularityBySummaryIdRequest) SetOpUserId(v string) *GetGranularityBySummaryIdRequest {
	s.OpUserId = &v
	return s
}

func (s *GetGranularityBySummaryIdRequest) SetProjectId(v int64) *GetGranularityBySummaryIdRequest {
	s.ProjectId = &v
	return s
}

func (s *GetGranularityBySummaryIdRequest) SetQueryDraft(v bool) *GetGranularityBySummaryIdRequest {
	s.QueryDraft = &v
	return s
}

func (s *GetGranularityBySummaryIdRequest) Validate() error {
	return dara.Validate(s)
}
