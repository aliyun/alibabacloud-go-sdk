// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMergeRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateMergeRequestRequest
	GetAccessToken() *string
	SetCreateFrom(v string) *CreateMergeRequestRequest
	GetCreateFrom() *string
	SetDescription(v string) *CreateMergeRequestRequest
	GetDescription() *string
	SetReviewerIds(v []*string) *CreateMergeRequestRequest
	GetReviewerIds() []*string
	SetSourceBranch(v string) *CreateMergeRequestRequest
	GetSourceBranch() *string
	SetSourceProjectId(v int64) *CreateMergeRequestRequest
	GetSourceProjectId() *int64
	SetTargetBranch(v string) *CreateMergeRequestRequest
	GetTargetBranch() *string
	SetTargetProjectId(v int64) *CreateMergeRequestRequest
	GetTargetProjectId() *int64
	SetTitle(v string) *CreateMergeRequestRequest
	GetTitle() *string
	SetWorkItemIds(v string) *CreateMergeRequestRequest
	GetWorkItemIds() *string
	SetOrganizationId(v string) *CreateMergeRequestRequest
	GetOrganizationId() *string
}

type CreateMergeRequestRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WEB
	CreateFrom  *string   `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	ReviewerIds []*string `json:"reviewerIds,omitempty" xml:"reviewerIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sourceBranch
	SourceBranch *string `json:"sourceBranch,omitempty" xml:"sourceBranch,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2369234
	SourceProjectId *int64 `json:"sourceProjectId,omitempty" xml:"sourceProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// targetBranch
	TargetBranch *string `json:"targetBranch,omitempty" xml:"targetBranch,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2369234
	TargetProjectId *int64 `json:"targetProjectId,omitempty" xml:"targetProjectId,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 722200214032b6b31e6f1434ab
	WorkItemIds *string `json:"workItemIds,omitempty" xml:"workItemIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6270e731cfea268afc21ccac
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateMergeRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateMergeRequestRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateMergeRequestRequest) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *CreateMergeRequestRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMergeRequestRequest) GetReviewerIds() []*string {
	return s.ReviewerIds
}

func (s *CreateMergeRequestRequest) GetSourceBranch() *string {
	return s.SourceBranch
}

func (s *CreateMergeRequestRequest) GetSourceProjectId() *int64 {
	return s.SourceProjectId
}

func (s *CreateMergeRequestRequest) GetTargetBranch() *string {
	return s.TargetBranch
}

func (s *CreateMergeRequestRequest) GetTargetProjectId() *int64 {
	return s.TargetProjectId
}

func (s *CreateMergeRequestRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateMergeRequestRequest) GetWorkItemIds() *string {
	return s.WorkItemIds
}

func (s *CreateMergeRequestRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateMergeRequestRequest) SetAccessToken(v string) *CreateMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateMergeRequestRequest) SetCreateFrom(v string) *CreateMergeRequestRequest {
	s.CreateFrom = &v
	return s
}

func (s *CreateMergeRequestRequest) SetDescription(v string) *CreateMergeRequestRequest {
	s.Description = &v
	return s
}

func (s *CreateMergeRequestRequest) SetReviewerIds(v []*string) *CreateMergeRequestRequest {
	s.ReviewerIds = v
	return s
}

func (s *CreateMergeRequestRequest) SetSourceBranch(v string) *CreateMergeRequestRequest {
	s.SourceBranch = &v
	return s
}

func (s *CreateMergeRequestRequest) SetSourceProjectId(v int64) *CreateMergeRequestRequest {
	s.SourceProjectId = &v
	return s
}

func (s *CreateMergeRequestRequest) SetTargetBranch(v string) *CreateMergeRequestRequest {
	s.TargetBranch = &v
	return s
}

func (s *CreateMergeRequestRequest) SetTargetProjectId(v int64) *CreateMergeRequestRequest {
	s.TargetProjectId = &v
	return s
}

func (s *CreateMergeRequestRequest) SetTitle(v string) *CreateMergeRequestRequest {
	s.Title = &v
	return s
}

func (s *CreateMergeRequestRequest) SetWorkItemIds(v string) *CreateMergeRequestRequest {
	s.WorkItemIds = &v
	return s
}

func (s *CreateMergeRequestRequest) SetOrganizationId(v string) *CreateMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateMergeRequestRequest) Validate() error {
	return dara.Validate(s)
}
