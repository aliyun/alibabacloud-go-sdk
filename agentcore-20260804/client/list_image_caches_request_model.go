// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *ListImageCachesRequest
	GetAcrInstanceId() *string
	SetDigest(v string) *ListImageCachesRequest
	GetDigest() *string
	SetImageRefContains(v string) *ListImageCachesRequest
	GetImageRefContains() *string
	SetMaxResults(v int32) *ListImageCachesRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListImageCachesRequest
	GetNamespace() *string
	SetNextToken(v string) *ListImageCachesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListImageCachesRequest
	GetRegionId() *string
	SetRegistry(v string) *ListImageCachesRequest
	GetRegistry() *string
	SetRepository(v string) *ListImageCachesRequest
	GetRepository() *string
	SetStatus(v string) *ListImageCachesRequest
	GetStatus() *string
	SetTag(v string) *ListImageCachesRequest
	GetTag() *string
	SetWorkspaceId(v string) *ListImageCachesRequest
	GetWorkspaceId() *string
}

type ListImageCachesRequest struct {
	// The ID of the source ACR Enterprise instance. The match must be exact. Empty character strings are invalid.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The complete image digest in the sha256: format followed by a 64-character lowercase hexadecimal string. The format is verified by the console. The match must be exact, and prefixes are not supported.
	//
	// example:
	//
	// sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	Digest *string `json:"digest,omitempty" xml:"digest,omitempty"`
	// Performs a case-insensitive literal substring search on the complete source image reference in the registry/repository:tag format, where the repository already includes the namespace. If the value is empty after trimming leading and trailing whitespace, it is treated as not provided. Percent signs, underscores, and backslashes are treated as literal characters.
	//
	// example:
	//
	// team/app
	ImageRefContains *string `json:"imageRefContains,omitempty" xml:"imageRefContains,omitempty"`
	// The maximum number of results to return. If this parameter is not provided or is set to 0, the default value is 20. If the value exceeds 100, it is treated as 100. Negative values return an InvalidMaxResults error. The value can be adjusted for subsequent pages.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The source namespace. The match is case-sensitive and must be exact. Empty strings are invalid.
	//
	// example:
	//
	// team
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// Omit this parameter for the first request. For subsequent requests, pass the nextToken value from the previous response exactly as returned. If the filter conditions change, you must query the first page again.
	//
	// example:
	//
	// opaque-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The product entry region, which is determined by the selected region endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The domain name of the source repository. The match is case-insensitive and must be exact. Do not include the protocol, path, or tag. Public network and VPC domain names are not considered equal. Empty strings are invalid.
	//
	// example:
	//
	// example-registry.cn-hangzhou.cr.aliyuncs.com
	Registry *string `json:"registry,omitempty" xml:"registry,omitempty"`
	// The name of the repository under the namespace. The match is case-sensitive and must be exact. Do not include the namespace. If this parameter is provided, the namespace parameter must also be provided. Empty strings are invalid.
	//
	// example:
	//
	// app
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// Filters the results by a single cache status. Omit this parameter to query all statuses.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the source image. The match is case-sensitive and must be exact. The latest tag is not inferred. Empty strings are invalid.
	//
	// example:
	//
	// v1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The current routed workspace. The server verifies the account, region, and resource ownership. This parameter cannot be omitted or used to access other workspaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListImageCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesRequest) GoString() string {
	return s.String()
}

func (s *ListImageCachesRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *ListImageCachesRequest) GetDigest() *string {
	return s.Digest
}

func (s *ListImageCachesRequest) GetImageRefContains() *string {
	return s.ImageRefContains
}

func (s *ListImageCachesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImageCachesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListImageCachesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImageCachesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImageCachesRequest) GetRegistry() *string {
	return s.Registry
}

func (s *ListImageCachesRequest) GetRepository() *string {
	return s.Repository
}

func (s *ListImageCachesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListImageCachesRequest) GetTag() *string {
	return s.Tag
}

func (s *ListImageCachesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListImageCachesRequest) SetAcrInstanceId(v string) *ListImageCachesRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *ListImageCachesRequest) SetDigest(v string) *ListImageCachesRequest {
	s.Digest = &v
	return s
}

func (s *ListImageCachesRequest) SetImageRefContains(v string) *ListImageCachesRequest {
	s.ImageRefContains = &v
	return s
}

func (s *ListImageCachesRequest) SetMaxResults(v int32) *ListImageCachesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListImageCachesRequest) SetNamespace(v string) *ListImageCachesRequest {
	s.Namespace = &v
	return s
}

func (s *ListImageCachesRequest) SetNextToken(v string) *ListImageCachesRequest {
	s.NextToken = &v
	return s
}

func (s *ListImageCachesRequest) SetRegionId(v string) *ListImageCachesRequest {
	s.RegionId = &v
	return s
}

func (s *ListImageCachesRequest) SetRegistry(v string) *ListImageCachesRequest {
	s.Registry = &v
	return s
}

func (s *ListImageCachesRequest) SetRepository(v string) *ListImageCachesRequest {
	s.Repository = &v
	return s
}

func (s *ListImageCachesRequest) SetStatus(v string) *ListImageCachesRequest {
	s.Status = &v
	return s
}

func (s *ListImageCachesRequest) SetTag(v string) *ListImageCachesRequest {
	s.Tag = &v
	return s
}

func (s *ListImageCachesRequest) SetWorkspaceId(v string) *ListImageCachesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListImageCachesRequest) Validate() error {
	return dara.Validate(s)
}
