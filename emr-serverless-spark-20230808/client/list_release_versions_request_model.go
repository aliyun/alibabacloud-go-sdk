// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReleaseVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListReleaseVersionsRequest
	GetRegionId() *string
	SetReleaseType(v string) *ListReleaseVersionsRequest
	GetReleaseType() *string
	SetReleaseVersion(v string) *ListReleaseVersionsRequest
	GetReleaseVersion() *string
	SetReleaseVersionStatus(v string) *ListReleaseVersionsRequest
	GetReleaseVersionStatus() *string
	SetServiceFilter(v string) *ListReleaseVersionsRequest
	GetServiceFilter() *string
	SetWorkspaceId(v string) *ListReleaseVersionsRequest
	GetWorkspaceId() *string
}

type ListReleaseVersionsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The type of the version.
	//
	// Valid values:
	//
	// 	- stable
	//
	// 	- Beta
	//
	// example:
	//
	// stable
	ReleaseType *string `json:"releaseType,omitempty" xml:"releaseType,omitempty"`
	// The version of EMR Serverless Spark.
	//
	// example:
	//
	// esr-2.1 (Spark 3.3.1, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The status of the version.
	//
	// Valid values:
	//
	// 	- ONLINE
	//
	// 	- OFFLINE
	//
	// example:
	//
	// ONLINE
	ReleaseVersionStatus *string `json:"releaseVersionStatus,omitempty" xml:"releaseVersionStatus,omitempty"`
	ServiceFilter        *string `json:"serviceFilter,omitempty" xml:"serviceFilter,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-d2d82aa09155****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListReleaseVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReleaseVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListReleaseVersionsRequest) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *ListReleaseVersionsRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListReleaseVersionsRequest) GetReleaseVersionStatus() *string {
	return s.ReleaseVersionStatus
}

func (s *ListReleaseVersionsRequest) GetServiceFilter() *string {
	return s.ServiceFilter
}

func (s *ListReleaseVersionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListReleaseVersionsRequest) SetRegionId(v string) *ListReleaseVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetReleaseType(v string) *ListReleaseVersionsRequest {
	s.ReleaseType = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetReleaseVersion(v string) *ListReleaseVersionsRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetReleaseVersionStatus(v string) *ListReleaseVersionsRequest {
	s.ReleaseVersionStatus = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetServiceFilter(v string) *ListReleaseVersionsRequest {
	s.ServiceFilter = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetWorkspaceId(v string) *ListReleaseVersionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListReleaseVersionsRequest) Validate() error {
	return dara.Validate(s)
}
