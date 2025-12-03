// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppReleaseStageExecutionIntegratedMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppReleaseStageExecutionIntegratedMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppReleaseStageExecutionIntegratedMetadataResponse
	GetStatusCode() *int32
	SetBody(v []*ListAppReleaseStageExecutionIntegratedMetadataResponseBody) *ListAppReleaseStageExecutionIntegratedMetadataResponse
	GetBody() []*ListAppReleaseStageExecutionIntegratedMetadataResponseBody
}

type ListAppReleaseStageExecutionIntegratedMetadataResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*ListAppReleaseStageExecutionIntegratedMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListAppReleaseStageExecutionIntegratedMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppReleaseStageExecutionIntegratedMetadataResponse) GoString() string {
	return s.String()
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponse) GetBody() []*ListAppReleaseStageExecutionIntegratedMetadataResponseBody {
	return s.Body
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponse) SetHeaders(v map[string]*string) *ListAppReleaseStageExecutionIntegratedMetadataResponse {
	s.Headers = v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponse) SetStatusCode(v int32) *ListAppReleaseStageExecutionIntegratedMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponse) SetBody(v []*ListAppReleaseStageExecutionIntegratedMetadataResponseBody) *ListAppReleaseStageExecutionIntegratedMetadataResponse {
	s.Body = v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponse) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppReleaseStageExecutionIntegratedMetadataResponseBody struct {
	// example:
	//
	// release/20240625-152603220321211_release_3252057_1
	ReleaseBranch *string `json:"releaseBranch,omitempty" xml:"releaseBranch,omitempty"`
	// example:
	//
	// a66cfa8c6869b96bb7d111ba2144c9c764d556b7
	ReleaseRevision *string `json:"releaseRevision,omitempty" xml:"releaseRevision,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/60d54f3daccf2bbd6659f3ad/auto-test.git
	RepoUrl *string `json:"repoUrl,omitempty" xml:"repoUrl,omitempty"`
	// example:
	//
	// CODEUP
	RepoType       *string                                                                     `json:"repoType,omitempty" xml:"repoType,omitempty"`
	ChangeRequests []*ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests `json:"changeRequests,omitempty" xml:"changeRequests,omitempty" type:"Repeated"`
}

func (s ListAppReleaseStageExecutionIntegratedMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppReleaseStageExecutionIntegratedMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) GetReleaseBranch() *string {
	return s.ReleaseBranch
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) GetReleaseRevision() *string {
	return s.ReleaseRevision
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) GetRepoUrl() *string {
	return s.RepoUrl
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) GetRepoType() *string {
	return s.RepoType
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) GetChangeRequests() []*ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests {
	return s.ChangeRequests
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) SetReleaseBranch(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBody {
	s.ReleaseBranch = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) SetReleaseRevision(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBody {
	s.ReleaseRevision = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) SetRepoUrl(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBody {
	s.RepoUrl = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) SetRepoType(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBody {
	s.RepoType = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) SetChangeRequests(v []*ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) *ListAppReleaseStageExecutionIntegratedMetadataResponseBody {
	s.ChangeRequests = v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBody) Validate() error {
	if s.ChangeRequests != nil {
		for _, item := range s.ChangeRequests {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests struct {
	// example:
	//
	// fcd37726a6f84c60b7eb9c5856007c2f
	Sn   *string `json:"sn,omitempty" xml:"sn,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// feature/20240625
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// example:
	//
	// a66cfa8c6869b96bb7d111ba2144c9c764d556b7
	CommitId *string `json:"commitId,omitempty" xml:"commitId,omitempty"`
	// example:
	//
	// 262579140573491041
	OwnerAccountId *string `json:"ownerAccountId,omitempty" xml:"ownerAccountId,omitempty"`
}

func (s ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) String() string {
	return dara.Prettify(s)
}

func (s ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) GoString() string {
	return s.String()
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) GetSn() *string {
	return s.Sn
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) GetName() *string {
	return s.Name
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) GetBranchName() *string {
	return s.BranchName
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) GetCommitId() *string {
	return s.CommitId
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) GetOwnerAccountId() *string {
	return s.OwnerAccountId
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) SetSn(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests {
	s.Sn = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) SetName(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests {
	s.Name = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) SetBranchName(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests {
	s.BranchName = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) SetCommitId(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests {
	s.CommitId = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) SetOwnerAccountId(v string) *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests {
	s.OwnerAccountId = &v
	return s
}

func (s *ListAppReleaseStageExecutionIntegratedMetadataResponseBodyChangeRequests) Validate() error {
	return dara.Validate(s)
}
