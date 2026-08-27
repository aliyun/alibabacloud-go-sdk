// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeLogs(v bool) *GetSkillRunRequest
	GetIncludeLogs() *bool
	SetRunId(v string) *GetSkillRunRequest
	GetRunId() *string
	SetTenantId(v string) *GetSkillRunRequest
	GetTenantId() *string
}

type GetSkillRunRequest struct {
	// Specifies whether to include execution logs. Default value: false. Enable this parameter only for troubleshooting.
	//
	// example:
	//
	// false
	IncludeLogs *bool `json:"includeLogs,omitempty" xml:"includeLogs,omitempty"`
	// The evaluation run ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleRunId
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The tenant ID to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSkillRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRunRequest) GoString() string {
	return s.String()
}

func (s *GetSkillRunRequest) GetIncludeLogs() *bool {
	return s.IncludeLogs
}

func (s *GetSkillRunRequest) GetRunId() *string {
	return s.RunId
}

func (s *GetSkillRunRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSkillRunRequest) SetIncludeLogs(v bool) *GetSkillRunRequest {
	s.IncludeLogs = &v
	return s
}

func (s *GetSkillRunRequest) SetRunId(v string) *GetSkillRunRequest {
	s.RunId = &v
	return s
}

func (s *GetSkillRunRequest) SetTenantId(v string) *GetSkillRunRequest {
	s.TenantId = &v
	return s
}

func (s *GetSkillRunRequest) Validate() error {
	return dara.Validate(s)
}
