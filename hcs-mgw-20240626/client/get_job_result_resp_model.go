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
	// The type of the data address created based on the files that failed to be migrated. This parameter is required if you create a data address.
	//
	// example:
	//
	// ossinv
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The number of files that are migrated. The number includes the number of files that are successfully migrated and the number of files that are skipped.
	//
	// example:
	//
	// 800
	CopiedObjectCount *int64 `json:"CopiedObjectCount,omitempty" xml:"CopiedObjectCount,omitempty"`
	// The data size of files that are migrated.
	//
	// example:
	//
	// 800
	CopiedObjectSize *int64 `json:"CopiedObjectSize,omitempty" xml:"CopiedObjectSize,omitempty"`
	// The number of files that failed to be migrated.
	//
	// example:
	//
	// 200
	FailedObjectCount *int64 `json:"FailedObjectCount,omitempty" xml:"FailedObjectCount,omitempty"`
	// The AccessKey ID that is used to access the bucket in which the inventory list of files that failed to be migrated resides. This parameter is required if you create a data address.
	//
	// example:
	//
	// test_access_id
	InvAccessId *string `json:"InvAccessId,omitempty" xml:"InvAccessId,omitempty"`
	// The AccessKey secret that is used to access the bucket in which the inventory list of files that failed to be migrated resides. This parameter is required if you create a data address.
	//
	// example:
	//
	// test_secret_key
	InvAccessSecret *string `json:"InvAccessSecret,omitempty" xml:"InvAccessSecret,omitempty"`
	// The name of the bucket in which the inventory list of files that failed to be migrated resides. This parameter is required if you create a data address.
	//
	// example:
	//
	// test_sys_bucket
	InvBucket *string `json:"InvBucket,omitempty" xml:"InvBucket,omitempty"`
	// The domain name of the bucket in which the inventory list of files that failed to be migrated resides. This parameter is required if you create a data address.
	//
	// example:
	//
	// test_domain
	InvDomain *string `json:"InvDomain,omitempty" xml:"InvDomain,omitempty"`
	// The type of the bucket in which the inventory list of files that failed to be migrated resides. This parameter is required if you create a data address.
	//
	// example:
	//
	// oss
	InvLocation *string `json:"InvLocation,omitempty" xml:"InvLocation,omitempty"`
	// The inventory list of files that failed to be migrated. This parameter is required if you create a data address.
	//
	// example:
	//
	// mainfest.json
	InvPath *string `json:"InvPath,omitempty" xml:"InvPath,omitempty"`
	// The region ID of the bucket in which the inventory list of files that failed to be migrated resides. This parameter is required if you create a data address.
	//
	// example:
	//
	// test_region_id
	InvRegionId *string `json:"InvRegionId,omitempty" xml:"InvRegionId,omitempty"`
	// Indicates whether the files that failed to be migrated can be migrated again.\\
	//
	// Valid values: NoNeed, Ready, and NotReady.
	//
	// example:
	//
	// Ready
	ReadyRetry *string `json:"ReadyRetry,omitempty" xml:"ReadyRetry,omitempty"`
	// example:
	//
	// 5000
	SkippedObjectCount *int64 `json:"SkippedObjectCount,omitempty" xml:"SkippedObjectCount,omitempty"`
	// example:
	//
	// 1000000
	SkippedObjectSize *int64 `json:"SkippedObjectSize,omitempty" xml:"SkippedObjectSize,omitempty"`
	// The total number of files.
	//
	// example:
	//
	// 1000
	TotalObjectCount *int64 `json:"TotalObjectCount,omitempty" xml:"TotalObjectCount,omitempty"`
	// The total data size of files.
	//
	// example:
	//
	// 1000
	TotalObjectSize *int64 `json:"TotalObjectSize,omitempty" xml:"TotalObjectSize,omitempty"`
	// The task ID.
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
