// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResultResp interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *GetJobResultResp
	GetAddressType() *string
	SetCopiedObjectCount(v int64) *GetJobResultResp
	GetCopiedObjectCount() *int64
	SetCopiedObjectSize(v int64) *GetJobResultResp
	GetCopiedObjectSize() *int64
	SetFailedObjectCount(v int64) *GetJobResultResp
	GetFailedObjectCount() *int64
	SetInvAccessId(v string) *GetJobResultResp
	GetInvAccessId() *string
	SetInvAccessSecret(v string) *GetJobResultResp
	GetInvAccessSecret() *string
	SetInvBucket(v string) *GetJobResultResp
	GetInvBucket() *string
	SetInvDomain(v string) *GetJobResultResp
	GetInvDomain() *string
	SetInvLocation(v string) *GetJobResultResp
	GetInvLocation() *string
	SetInvPath(v string) *GetJobResultResp
	GetInvPath() *string
	SetInvRegionId(v string) *GetJobResultResp
	GetInvRegionId() *string
	SetReadyRetry(v string) *GetJobResultResp
	GetReadyRetry() *string
	SetSkippedObjectCount(v int64) *GetJobResultResp
	GetSkippedObjectCount() *int64
	SetSkippedObjectSize(v int64) *GetJobResultResp
	GetSkippedObjectSize() *int64
	SetTotalObjectCount(v int64) *GetJobResultResp
	GetTotalObjectCount() *int64
	SetTotalObjectSize(v int64) *GetJobResultResp
	GetTotalObjectSize() *int64
	SetVersion(v string) *GetJobResultResp
	GetVersion() *string
}

type GetJobResultResp struct {
	// The data address type for the retry job. This value indicates that the data address is constructed from a failed file inventory. Use this value as the AddressType parameter when you create a data address for a retry job.
	//
	// example:
	//
	// ossinv
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The number of objects that were processed successfully. This value includes both migrated objects and skipped objects.
	//
	// example:
	//
	// 800
	CopiedObjectCount *int64 `json:"CopiedObjectCount,omitempty" xml:"CopiedObjectCount,omitempty"`
	// The total size of objects that were processed successfully. Unit: bytes.
	//
	// example:
	//
	// 800
	CopiedObjectSize *int64 `json:"CopiedObjectSize,omitempty" xml:"CopiedObjectSize,omitempty"`
	// The number of objects that failed to migrate.
	//
	// example:
	//
	// 200
	FailedObjectCount *int64 `json:"FailedObjectCount,omitempty" xml:"FailedObjectCount,omitempty"`
	// The AccessKey ID that is used to access the bucket where the failed file list is stored. Use this value as the InvAccessId parameter when you create a data address for a retry job.
	//
	// example:
	//
	// test_access_id
	InvAccessId *string `json:"InvAccessId,omitempty" xml:"InvAccessId,omitempty"`
	// The AccessKey secret that is used to access the bucket where the failed file list is stored. Use this value as the InvAccessSecret parameter when you create a data address for a retry job.
	//
	// example:
	//
	// test_secret_key
	InvAccessSecret *string `json:"InvAccessSecret,omitempty" xml:"InvAccessSecret,omitempty"`
	// The name of the bucket that stores the failed file list. Use this value as the InvBucket parameter when you create a data address for a retry job.
	//
	// example:
	//
	// test_sys_bucket
	InvBucket *string `json:"InvBucket,omitempty" xml:"InvBucket,omitempty"`
	// The endpoint of the bucket that stores the failed file list. Use this value as the InvDomain parameter when you create a data address for a retry job.
	//
	// example:
	//
	// test_domain
	InvDomain *string `json:"InvDomain,omitempty" xml:"InvDomain,omitempty"`
	// The storage type of the bucket that stores the failed file list, such as oss. Use this value as the InvLocation parameter when you create a data address for a retry job.
	//
	// example:
	//
	// oss
	InvLocation *string `json:"InvLocation,omitempty" xml:"InvLocation,omitempty"`
	// The path to the manifest file that lists the failed files. Use this value as the InvPath parameter when you create a data address for a retry job.
	//
	// example:
	//
	// mainfest.json
	InvPath *string `json:"InvPath,omitempty" xml:"InvPath,omitempty"`
	// The region ID of the bucket that stores the failed file list. Use this value as the InvRegionId parameter when you create a data address for a retry job.
	//
	// example:
	//
	// test_region_id
	InvRegionId *string `json:"InvRegionId,omitempty" xml:"InvRegionId,omitempty"`
	// The retry readiness status for failed files. Valid values: NoNeed indicates that all files were migrated successfully and no retry is required. Ready indicates that the failed file list has been generated and is available for retry. NotReady indicates that the failed file list is being generated.<br><br>
	//
	// example:
	//
	// Ready
	ReadyRetry *string `json:"ReadyRetry,omitempty" xml:"ReadyRetry,omitempty"`
	// The number of objects that were skipped during migration. Objects are skipped when they already exist at the destination and meet the configured skip conditions.
	//
	// example:
	//
	// 5000
	SkippedObjectCount *int64 `json:"SkippedObjectCount,omitempty" xml:"SkippedObjectCount,omitempty"`
	// The total size of objects that were skipped during migration. Unit: bytes.
	//
	// example:
	//
	// 1000000
	SkippedObjectSize *int64 `json:"SkippedObjectSize,omitempty" xml:"SkippedObjectSize,omitempty"`
	// The total number of objects in the source data address.
	//
	// example:
	//
	// 1000
	TotalObjectCount *int64 `json:"TotalObjectCount,omitempty" xml:"TotalObjectCount,omitempty"`
	// The total size of all objects in the source data address. Unit: bytes.
	//
	// example:
	//
	// 1000
	TotalObjectSize *int64 `json:"TotalObjectSize,omitempty" xml:"TotalObjectSize,omitempty"`
	// The unique identifier of the migration job.
	//
	// example:
	//
	// test_job_id
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetJobResultResp) String() string {
	return dara.Prettify(s)
}

func (s GetJobResultResp) GoString() string {
	return s.String()
}

func (s *GetJobResultResp) GetAddressType() *string {
	return s.AddressType
}

func (s *GetJobResultResp) GetCopiedObjectCount() *int64 {
	return s.CopiedObjectCount
}

func (s *GetJobResultResp) GetCopiedObjectSize() *int64 {
	return s.CopiedObjectSize
}

func (s *GetJobResultResp) GetFailedObjectCount() *int64 {
	return s.FailedObjectCount
}

func (s *GetJobResultResp) GetInvAccessId() *string {
	return s.InvAccessId
}

func (s *GetJobResultResp) GetInvAccessSecret() *string {
	return s.InvAccessSecret
}

func (s *GetJobResultResp) GetInvBucket() *string {
	return s.InvBucket
}

func (s *GetJobResultResp) GetInvDomain() *string {
	return s.InvDomain
}

func (s *GetJobResultResp) GetInvLocation() *string {
	return s.InvLocation
}

func (s *GetJobResultResp) GetInvPath() *string {
	return s.InvPath
}

func (s *GetJobResultResp) GetInvRegionId() *string {
	return s.InvRegionId
}

func (s *GetJobResultResp) GetReadyRetry() *string {
	return s.ReadyRetry
}

func (s *GetJobResultResp) GetSkippedObjectCount() *int64 {
	return s.SkippedObjectCount
}

func (s *GetJobResultResp) GetSkippedObjectSize() *int64 {
	return s.SkippedObjectSize
}

func (s *GetJobResultResp) GetTotalObjectCount() *int64 {
	return s.TotalObjectCount
}

func (s *GetJobResultResp) GetTotalObjectSize() *int64 {
	return s.TotalObjectSize
}

func (s *GetJobResultResp) GetVersion() *string {
	return s.Version
}

func (s *GetJobResultResp) SetAddressType(v string) *GetJobResultResp {
	s.AddressType = &v
	return s
}

func (s *GetJobResultResp) SetCopiedObjectCount(v int64) *GetJobResultResp {
	s.CopiedObjectCount = &v
	return s
}

func (s *GetJobResultResp) SetCopiedObjectSize(v int64) *GetJobResultResp {
	s.CopiedObjectSize = &v
	return s
}

func (s *GetJobResultResp) SetFailedObjectCount(v int64) *GetJobResultResp {
	s.FailedObjectCount = &v
	return s
}

func (s *GetJobResultResp) SetInvAccessId(v string) *GetJobResultResp {
	s.InvAccessId = &v
	return s
}

func (s *GetJobResultResp) SetInvAccessSecret(v string) *GetJobResultResp {
	s.InvAccessSecret = &v
	return s
}

func (s *GetJobResultResp) SetInvBucket(v string) *GetJobResultResp {
	s.InvBucket = &v
	return s
}

func (s *GetJobResultResp) SetInvDomain(v string) *GetJobResultResp {
	s.InvDomain = &v
	return s
}

func (s *GetJobResultResp) SetInvLocation(v string) *GetJobResultResp {
	s.InvLocation = &v
	return s
}

func (s *GetJobResultResp) SetInvPath(v string) *GetJobResultResp {
	s.InvPath = &v
	return s
}

func (s *GetJobResultResp) SetInvRegionId(v string) *GetJobResultResp {
	s.InvRegionId = &v
	return s
}

func (s *GetJobResultResp) SetReadyRetry(v string) *GetJobResultResp {
	s.ReadyRetry = &v
	return s
}

func (s *GetJobResultResp) SetSkippedObjectCount(v int64) *GetJobResultResp {
	s.SkippedObjectCount = &v
	return s
}

func (s *GetJobResultResp) SetSkippedObjectSize(v int64) *GetJobResultResp {
	s.SkippedObjectSize = &v
	return s
}

func (s *GetJobResultResp) SetTotalObjectCount(v int64) *GetJobResultResp {
	s.TotalObjectCount = &v
	return s
}

func (s *GetJobResultResp) SetTotalObjectSize(v int64) *GetJobResultResp {
	s.TotalObjectSize = &v
	return s
}

func (s *GetJobResultResp) SetVersion(v string) *GetJobResultResp {
	s.Version = &v
	return s
}

func (s *GetJobResultResp) Validate() error {
	return dara.Validate(s)
}
