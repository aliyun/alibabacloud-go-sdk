// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectedResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedByProduct(v string) *ListProtectedResourcesRequest
	GetCreatedByProduct() *string
	SetHasSnapshot(v bool) *ListProtectedResourcesRequest
	GetHasSnapshot() *bool
	SetMaxResults(v int32) *ListProtectedResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProtectedResourcesRequest
	GetNextToken() *string
	SetResourceId(v string) *ListProtectedResourcesRequest
	GetResourceId() *string
	SetSkip(v int32) *ListProtectedResourcesRequest
	GetSkip() *int32
	SetSourceType(v string) *ListProtectedResourcesRequest
	GetSourceType() *string
}

type ListProtectedResourcesRequest struct {
	// The product capability to which the resource belongs. Valid values:
	//
	// - **HBR**: Cloud Backup standard capability.
	//
	// - **BASIC**: ECS File Backup Essential Edition.
	//
	// example:
	//
	// BASIC
	CreatedByProduct *string `json:"CreatedByProduct,omitempty" xml:"CreatedByProduct,omitempty"`
	// Specifies whether the resource has backup points.
	//
	// example:
	//
	// true
	HasSnapshot *bool `json:"HasSnapshot,omitempty" xml:"HasSnapshot,omitempty"`
	// The number of results per query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. If this parameter is empty, no more pages are available.
	//
	// example:
	//
	// aWQj********MCMy
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource ID.
	//
	// - **SourceType=ECS_FILE**: the ECS instance ID.
	//
	// - **SourceType=COMMON_FILE_SYSTEM**: the CPFS data source ID.
	//
	// - **SourceType=COMMON_NAS**: the on-premises NAS data source ID.
	//
	// - **SourceType=File**: the local service client ID.
	//
	// - **SourceType=NAS**: the Alibaba Cloud NAS file system ID.
	//
	// - **SourceType=OSS**: the OSS bucket.
	//
	// example:
	//
	// i-wz95************7zrd
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The number of entries to skip for paging.
	//
	// If the number of skipped entries exceeds the total number of conditional entries, an empty list is returned. The number of skipped entries must be a multiple of MaxResults.
	//
	// example:
	//
	// 10
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The backup feature type. Valid values:
	//
	// - **ECS_FILE**: ECS file backup.
	//
	// - **COMMON_FILE_SYSTEM**: Cloud Parallel File Storage (CPFS) backup.
	//
	// - **COMMON_NAS**: on-premises NAS backup.
	//
	// - **File**: on-premises file backup.
	//
	// - **NAS**: Alibaba Cloud NAS backup.
	//
	// - **OSS**: OSS backup.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListProtectedResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListProtectedResourcesRequest) GetCreatedByProduct() *string {
	return s.CreatedByProduct
}

func (s *ListProtectedResourcesRequest) GetHasSnapshot() *bool {
	return s.HasSnapshot
}

func (s *ListProtectedResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProtectedResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProtectedResourcesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListProtectedResourcesRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProtectedResourcesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListProtectedResourcesRequest) SetCreatedByProduct(v string) *ListProtectedResourcesRequest {
	s.CreatedByProduct = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetHasSnapshot(v bool) *ListProtectedResourcesRequest {
	s.HasSnapshot = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetMaxResults(v int32) *ListProtectedResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetNextToken(v string) *ListProtectedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetResourceId(v string) *ListProtectedResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetSkip(v int32) *ListProtectedResourcesRequest {
	s.Skip = &v
	return s
}

func (s *ListProtectedResourcesRequest) SetSourceType(v string) *ListProtectedResourcesRequest {
	s.SourceType = &v
	return s
}

func (s *ListProtectedResourcesRequest) Validate() error {
	return dara.Validate(s)
}
