// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrossProjectDeploymentCandidatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeType(v string) *ListCrossProjectDeploymentCandidatesRequest
	GetChangeType() *string
	SetCommitTimeFrom(v int64) *ListCrossProjectDeploymentCandidatesRequest
	GetCommitTimeFrom() *int64
	SetCommitTimeTo(v int64) *ListCrossProjectDeploymentCandidatesRequest
	GetCommitTimeTo() *int64
	SetCommitUser(v string) *ListCrossProjectDeploymentCandidatesRequest
	GetCommitUser() *string
	SetDeploymentEnvironmentId(v int64) *ListCrossProjectDeploymentCandidatesRequest
	GetDeploymentEnvironmentId() *int64
	SetKeyword(v string) *ListCrossProjectDeploymentCandidatesRequest
	GetKeyword() *string
	SetObjectId(v string) *ListCrossProjectDeploymentCandidatesRequest
	GetObjectId() *string
	SetObjectType(v string) *ListCrossProjectDeploymentCandidatesRequest
	GetObjectType() *string
	SetPageNumber(v int32) *ListCrossProjectDeploymentCandidatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCrossProjectDeploymentCandidatesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCrossProjectDeploymentCandidatesRequest
	GetProjectId() *int64
}

type ListCrossProjectDeploymentCandidatesRequest struct {
	// The change type.
	//
	// example:
	//
	// ADD
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The start of the commit time range. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788739200000
	CommitTimeFrom *int64 `json:"CommitTimeFrom,omitempty" xml:"CommitTimeFrom,omitempty"`
	// The end of the commit time range. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788825599999
	CommitTimeTo *int64 `json:"CommitTimeTo,omitempty" xml:"CommitTimeTo,omitempty"`
	// The committer.
	//
	// example:
	//
	// operator
	CommitUser *string `json:"CommitUser,omitempty" xml:"CommitUser,omitempty"`
	// The cross-workspace deployment environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 101
	DeploymentEnvironmentId *int64 `json:"DeploymentEnvironmentId,omitempty" xml:"DeploymentEnvironmentId,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// object
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The candidate object ID.
	//
	// example:
	//
	// 1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The candidate object type.
	//
	// example:
	//
	// ODPS_SQL
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListCrossProjectDeploymentCandidatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrossProjectDeploymentCandidatesRequest) GoString() string {
	return s.String()
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetChangeType() *string {
	return s.ChangeType
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetCommitTimeFrom() *int64 {
	return s.CommitTimeFrom
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetCommitTimeTo() *int64 {
	return s.CommitTimeTo
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetCommitUser() *string {
	return s.CommitUser
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetDeploymentEnvironmentId() *int64 {
	return s.DeploymentEnvironmentId
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrossProjectDeploymentCandidatesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetChangeType(v string) *ListCrossProjectDeploymentCandidatesRequest {
	s.ChangeType = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetCommitTimeFrom(v int64) *ListCrossProjectDeploymentCandidatesRequest {
	s.CommitTimeFrom = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetCommitTimeTo(v int64) *ListCrossProjectDeploymentCandidatesRequest {
	s.CommitTimeTo = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetCommitUser(v string) *ListCrossProjectDeploymentCandidatesRequest {
	s.CommitUser = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetDeploymentEnvironmentId(v int64) *ListCrossProjectDeploymentCandidatesRequest {
	s.DeploymentEnvironmentId = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetKeyword(v string) *ListCrossProjectDeploymentCandidatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetObjectId(v string) *ListCrossProjectDeploymentCandidatesRequest {
	s.ObjectId = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetObjectType(v string) *ListCrossProjectDeploymentCandidatesRequest {
	s.ObjectType = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetPageNumber(v int32) *ListCrossProjectDeploymentCandidatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetPageSize(v int32) *ListCrossProjectDeploymentCandidatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) SetProjectId(v int64) *ListCrossProjectDeploymentCandidatesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCrossProjectDeploymentCandidatesRequest) Validate() error {
	return dara.Validate(s)
}
