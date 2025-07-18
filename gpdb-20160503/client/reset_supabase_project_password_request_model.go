// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSupabaseProjectPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountPassword(v string) *ResetSupabaseProjectPasswordRequest
	GetAccountPassword() *string
	SetProjectId(v string) *ResetSupabaseProjectPasswordRequest
	GetProjectId() *string
	SetRegionId(v string) *ResetSupabaseProjectPasswordRequest
	GetRegionId() *string
}

type ResetSupabaseProjectPasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Pw123456
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sbp-tyarplz****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetSupabaseProjectPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetSupabaseProjectPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetSupabaseProjectPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetSupabaseProjectPasswordRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ResetSupabaseProjectPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetSupabaseProjectPasswordRequest) SetAccountPassword(v string) *ResetSupabaseProjectPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetSupabaseProjectPasswordRequest) SetProjectId(v string) *ResetSupabaseProjectPasswordRequest {
	s.ProjectId = &v
	return s
}

func (s *ResetSupabaseProjectPasswordRequest) SetRegionId(v string) *ResetSupabaseProjectPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSupabaseProjectPasswordRequest) Validate() error {
	return dara.Validate(s)
}
