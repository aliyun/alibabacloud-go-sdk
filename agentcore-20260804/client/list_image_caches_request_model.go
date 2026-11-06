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
	// The ID of the source ACR Enterprise instance. The match must be exact. An empty string (a string with no characters) is invalid.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The image digest.
	//
	// example:
	//
	// sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	Digest *string `json:"digest,omitempty" xml:"digest,omitempty"`
	// Performs a case-insensitive literal substring search on the full source image reference in the format of registry/repository:tag. The repository already includes the namespace. If the value is empty after trimming leading and trailing whitespace, it is treated as not provided. Percent signs, underscores, and backslashes are treated as literal characters.
	//
	// example:
	//
	// team/app
	ImageRefContains *string `json:"imageRefContains,omitempty" xml:"imageRefContains,omitempty"`
	// The maximum number of results to return. Default value: 20. If you explicitly specify this parameter, the value must be an integer from 1 to 100. Otherwise, an InvalidMaxResults error is returned. When you paginate, this value must remain the same as the previous page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The source namespace. The match is case-sensitive and must be exact. An empty string is invalid.
	//
	// example:
	//
	// team
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The pagination token. Omit this parameter for the first request. For subsequent requests, pass the nextToken value returned in the previous response as is. When you paginate, you must keep the filter conditions and maxResults unchanged. If you change any of these parameters, you must query the first page again.
	//
	// example:
	//
	// opaque-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region ID. The region ID must be passed as a query parameter and must match the selected region endpoint and the workspace region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The domain name of the source registry. The match is case-insensitive and must be exact. Do not include the protocol, path, or tag. Public network and VPC domain names are not considered equal. An empty string is invalid.
	//
	// example:
	//
	// example-registry.cn-hangzhou.cr.aliyuncs.com
	Registry *string `json:"registry,omitempty" xml:"registry,omitempty"`
	// The repository name under the namespace. The match is case-sensitive and must be exact. Do not include the namespace. If you specify this parameter, you must also specify the namespace. An empty string is invalid.
	//
	// example:
	//
	// app
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// The cache status used to filter the results. To query all caches, omit this parameter.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The source image tag. The match is case-sensitive and must be exact. The latest tag is not inferred. An empty string is invalid.
	//
	// example:
	//
	// v1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The workspace ID.
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
