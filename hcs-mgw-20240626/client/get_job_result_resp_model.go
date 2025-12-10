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
	// example:
	//
	// ossinv
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// 800
	CopiedObjectCount *int64 `json:"CopiedObjectCount,omitempty" xml:"CopiedObjectCount,omitempty"`
	// example:
	//
	// 800
	CopiedObjectSize *int64 `json:"CopiedObjectSize,omitempty" xml:"CopiedObjectSize,omitempty"`
	// example:
	//
	// 200
	FailedObjectCount *int64 `json:"FailedObjectCount,omitempty" xml:"FailedObjectCount,omitempty"`
	// example:
	//
	// **********************
	InvAccessId *string `json:"InvAccessId,omitempty" xml:"InvAccessId,omitempty"`
	// example:
	//
	// *************************
	InvAccessSecret *string `json:"InvAccessSecret,omitempty" xml:"InvAccessSecret,omitempty"`
	// example:
	//
	// <your-bucket-name>
	InvBucket *string `json:"InvBucket,omitempty" xml:"InvBucket,omitempty"`
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	InvDomain *string `json:"InvDomain,omitempty" xml:"InvDomain,omitempty"`
	// example:
	//
	// oss
	InvLocation *string `json:"InvLocation,omitempty" xml:"InvLocation,omitempty"`
	// example:
	//
	// dir/mainfest.json
	InvPath *string `json:"InvPath,omitempty" xml:"InvPath,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	InvRegionId *string `json:"InvRegionId,omitempty" xml:"InvRegionId,omitempty"`
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
	// example:
	//
	// 1000
	TotalObjectCount *int64 `json:"TotalObjectCount,omitempty" xml:"TotalObjectCount,omitempty"`
	// example:
	//
	// 1000
	TotalObjectSize *int64 `json:"TotalObjectSize,omitempty" xml:"TotalObjectSize,omitempty"`
	// example:
	//
	// ******-188f-41d9-b266-******
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
