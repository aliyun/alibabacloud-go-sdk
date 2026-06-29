// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *RestoreBranchRequest
	GetBranchId() *string
	SetClientToken(v string) *RestoreBranchRequest
	GetClientToken() *string
	SetPreserveUnderName(v string) *RestoreBranchRequest
	GetPreserveUnderName() *string
	SetProjectId(v string) *RestoreBranchRequest
	GetProjectId() *string
	SetRegionId(v string) *RestoreBranchRequest
	GetRegionId() *string
	SetSourceBranchId(v string) *RestoreBranchRequest
	GetSourceBranchId() *string
	SetSourceBranchLsn(v string) *RestoreBranchRequest
	GetSourceBranchLsn() *string
	SetSourceBranchTimestamp(v string) *RestoreBranchRequest
	GetSourceBranchTimestamp() *string
}

type RestoreBranchRequest struct {
	// The branch ID that uniquely identifies a Supabase branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The client idempotency token that ensures the idempotence of retry requests.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The backup branch name. If specified, automatic creation of a backup branch is performed before recovery.
	//
	// example:
	//
	// backup-main
	PreserveUnderName *string `json:"PreserveUnderName,omitempty" xml:"PreserveUnderName,omitempty"`
	// The Supabase project ID associated with the primary branch.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID. This parameter is required when you create a primary branch. When you create a sub-branch, the region is inherited from the primary branch by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the source branch from which to recover.
	//
	// This parameter is required.
	//
	// example:
	//
	// br-main
	SourceBranchId *string `json:"SourceBranchId,omitempty" xml:"SourceBranchId,omitempty"`
	// The LSN of the source branch to recover to.
	//
	// example:
	//
	// 0/3522648
	SourceBranchLsn *string `json:"SourceBranchLsn,omitempty" xml:"SourceBranchLsn,omitempty"`
	// The point in time of the source branch to recover to. The value must be within the recoverable time window.
	//
	// example:
	//
	// 2026-04-08T09:11:12Z
	SourceBranchTimestamp *string `json:"SourceBranchTimestamp,omitempty" xml:"SourceBranchTimestamp,omitempty"`
}

func (s RestoreBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreBranchRequest) GoString() string {
	return s.String()
}

func (s *RestoreBranchRequest) GetBranchId() *string {
	return s.BranchId
}

func (s *RestoreBranchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestoreBranchRequest) GetPreserveUnderName() *string {
	return s.PreserveUnderName
}

func (s *RestoreBranchRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *RestoreBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestoreBranchRequest) GetSourceBranchId() *string {
	return s.SourceBranchId
}

func (s *RestoreBranchRequest) GetSourceBranchLsn() *string {
	return s.SourceBranchLsn
}

func (s *RestoreBranchRequest) GetSourceBranchTimestamp() *string {
	return s.SourceBranchTimestamp
}

func (s *RestoreBranchRequest) SetBranchId(v string) *RestoreBranchRequest {
	s.BranchId = &v
	return s
}

func (s *RestoreBranchRequest) SetClientToken(v string) *RestoreBranchRequest {
	s.ClientToken = &v
	return s
}

func (s *RestoreBranchRequest) SetPreserveUnderName(v string) *RestoreBranchRequest {
	s.PreserveUnderName = &v
	return s
}

func (s *RestoreBranchRequest) SetProjectId(v string) *RestoreBranchRequest {
	s.ProjectId = &v
	return s
}

func (s *RestoreBranchRequest) SetRegionId(v string) *RestoreBranchRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreBranchRequest) SetSourceBranchId(v string) *RestoreBranchRequest {
	s.SourceBranchId = &v
	return s
}

func (s *RestoreBranchRequest) SetSourceBranchLsn(v string) *RestoreBranchRequest {
	s.SourceBranchLsn = &v
	return s
}

func (s *RestoreBranchRequest) SetSourceBranchTimestamp(v string) *RestoreBranchRequest {
	s.SourceBranchTimestamp = &v
	return s
}

func (s *RestoreBranchRequest) Validate() error {
	return dara.Validate(s)
}
