// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectedResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProtectedResourcesResponseBody
	GetCode() *string
	SetMaxResults(v int32) *ListProtectedResourcesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListProtectedResourcesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListProtectedResourcesResponseBody
	GetNextToken() *string
	SetProtectedResources(v []*ListProtectedResourcesResponseBodyProtectedResources) *ListProtectedResourcesResponseBody
	GetProtectedResources() []*ListProtectedResourcesResponseBodyProtectedResources
	SetRequestId(v string) *ListProtectedResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProtectedResourcesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListProtectedResourcesResponseBody
	GetTotalCount() *int32
}

type ListProtectedResourcesResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of results per query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message. The value "successful" is returned for a successful request. An error message is returned for a failed request.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token for the next page. If this parameter is empty, no more pages are available.
	//
	// example:
	//
	// eyJJ************MX0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of protected resources.
	ProtectedResources []*ListProtectedResourcesResponseBodyProtectedResources `json:"ProtectedResources,omitempty" xml:"ProtectedResources,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EB09****-****-****-****-********6C38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of protected resources.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProtectedResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProtectedResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProtectedResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProtectedResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProtectedResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProtectedResourcesResponseBody) GetProtectedResources() []*ListProtectedResourcesResponseBodyProtectedResources {
	return s.ProtectedResources
}

func (s *ListProtectedResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProtectedResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProtectedResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProtectedResourcesResponseBody) SetCode(v string) *ListProtectedResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetMaxResults(v int32) *ListProtectedResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetMessage(v string) *ListProtectedResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetNextToken(v string) *ListProtectedResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetProtectedResources(v []*ListProtectedResourcesResponseBodyProtectedResources) *ListProtectedResourcesResponseBody {
	s.ProtectedResources = v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetRequestId(v string) *ListProtectedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetSuccess(v bool) *ListProtectedResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) SetTotalCount(v int32) *ListProtectedResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProtectedResourcesResponseBody) Validate() error {
	if s.ProtectedResources != nil {
		for _, item := range s.ProtectedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProtectedResourcesResponseBodyProtectedResources struct {
	// The number of backup plans.
	//
	// example:
	//
	// 1
	BackupPlanCount *int64 `json:"BackupPlanCount,omitempty" xml:"BackupPlanCount,omitempty"`
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
	// The amount of protected data, in bytes. Currently, only ECS File Backup Essential Edition is supported.
	//
	// - **SourceType=ECS_FILE**: the backed-up block storage capacity.
	//
	// example:
	//
	// 107374182400
	ProtectedDataSize *int64 `json:"ProtectedDataSize,omitempty" xml:"ProtectedDataSize,omitempty"`
	// The ID of the protected resource.
	//
	// example:
	//
	// pr-0004************gs61
	ProtectedResourceId *string `json:"ProtectedResourceId,omitempty" xml:"ProtectedResourceId,omitempty"`
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
	// The UID of the user who owns the resource.
	//
	// example:
	//
	// 1024********0703
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The number of backups.
	//
	// example:
	//
	// 30
	SnapshotCount *int64 `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
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

func (s ListProtectedResourcesResponseBodyProtectedResources) String() string {
	return dara.Prettify(s)
}

func (s ListProtectedResourcesResponseBodyProtectedResources) GoString() string {
	return s.String()
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetBackupPlanCount() *int64 {
	return s.BackupPlanCount
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetCreatedByProduct() *string {
	return s.CreatedByProduct
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetProtectedDataSize() *int64 {
	return s.ProtectedDataSize
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetProtectedResourceId() *string {
	return s.ProtectedResourceId
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetSnapshotCount() *int64 {
	return s.SnapshotCount
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) GetSourceType() *string {
	return s.SourceType
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetBackupPlanCount(v int64) *ListProtectedResourcesResponseBodyProtectedResources {
	s.BackupPlanCount = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetCreatedByProduct(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.CreatedByProduct = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetProtectedDataSize(v int64) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ProtectedDataSize = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetProtectedResourceId(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ProtectedResourceId = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetResourceId(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ResourceId = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetResourceOwnerId(v int64) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetResourceRegionId(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.ResourceRegionId = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetSnapshotCount(v int64) *ListProtectedResourcesResponseBodyProtectedResources {
	s.SnapshotCount = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) SetSourceType(v string) *ListProtectedResourcesResponseBodyProtectedResources {
	s.SourceType = &v
	return s
}

func (s *ListProtectedResourcesResponseBodyProtectedResources) Validate() error {
	return dara.Validate(s)
}
