// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupStorageCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupStorageCount(v *GetBackupStorageCountResponseBodyBackupStorageCount) *GetBackupStorageCountResponseBody
	GetBackupStorageCount() *GetBackupStorageCountResponseBodyBackupStorageCount
	SetRequestId(v string) *GetBackupStorageCountResponseBody
	GetRequestId() *string
}

type GetBackupStorageCountResponseBody struct {
	// The details about the anti-ransomware capacity.
	BackupStorageCount *GetBackupStorageCountResponseBodyBackupStorageCount `json:"BackupStorageCount,omitempty" xml:"BackupStorageCount,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 33C2CCFF-4BF8-5F88-9B5C-22F932F80E5A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBackupStorageCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBackupStorageCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupStorageCountResponseBody) GetBackupStorageCount() *GetBackupStorageCountResponseBodyBackupStorageCount {
	return s.BackupStorageCount
}

func (s *GetBackupStorageCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBackupStorageCountResponseBody) SetBackupStorageCount(v *GetBackupStorageCountResponseBodyBackupStorageCount) *GetBackupStorageCountResponseBody {
	s.BackupStorageCount = v
	return s
}

func (s *GetBackupStorageCountResponseBody) SetRequestId(v string) *GetBackupStorageCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackupStorageCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBackupStorageCountResponseBodyBackupStorageCount struct {
	// The anti-ransomware capacity that you purchased. Unit: bytes.
	//
	// example:
	//
	// 2276332666880
	BuyStorageByte *int64 `json:"BuyStorageByte,omitempty" xml:"BuyStorageByte,omitempty"`
	// The storage capacity that is occupied by the backup data of your servers. Unit: bytes.
	//
	// example:
	//
	// 817262417803
	EcsUsageStorageByte *int64 `json:"EcsUsageStorageByte,omitempty" xml:"EcsUsageStorageByte,omitempty"`
	// Indicates whether the anti-ransomware capacity that is used exceeds the anti-ransomware capacity that you purchased. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	Overflow *int32 `json:"Overflow,omitempty" xml:"Overflow,omitempty"`
	// The storage capacity that is occupied by the backup data of your databases. Unit: bytes.
	//
	// example:
	//
	// 7453049350
	UniUsageStorageByte *int64 `json:"UniUsageStorageByte,omitempty" xml:"UniUsageStorageByte,omitempty"`
	// The total anti-ransomware capacity that is used. Unit: bytes.
	//
	// example:
	//
	// 839621565853
	UsageStorageByte *int64 `json:"UsageStorageByte,omitempty" xml:"UsageStorageByte,omitempty"`
}

func (s GetBackupStorageCountResponseBodyBackupStorageCount) String() string {
	return dara.Prettify(s)
}

func (s GetBackupStorageCountResponseBodyBackupStorageCount) GoString() string {
	return s.String()
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) GetBuyStorageByte() *int64 {
	return s.BuyStorageByte
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) GetEcsUsageStorageByte() *int64 {
	return s.EcsUsageStorageByte
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) GetOverflow() *int32 {
	return s.Overflow
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) GetUniUsageStorageByte() *int64 {
	return s.UniUsageStorageByte
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) GetUsageStorageByte() *int64 {
	return s.UsageStorageByte
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) SetBuyStorageByte(v int64) *GetBackupStorageCountResponseBodyBackupStorageCount {
	s.BuyStorageByte = &v
	return s
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) SetEcsUsageStorageByte(v int64) *GetBackupStorageCountResponseBodyBackupStorageCount {
	s.EcsUsageStorageByte = &v
	return s
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) SetOverflow(v int32) *GetBackupStorageCountResponseBodyBackupStorageCount {
	s.Overflow = &v
	return s
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) SetUniUsageStorageByte(v int64) *GetBackupStorageCountResponseBodyBackupStorageCount {
	s.UniUsageStorageByte = &v
	return s
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) SetUsageStorageByte(v int64) *GetBackupStorageCountResponseBodyBackupStorageCount {
	s.UsageStorageByte = &v
	return s
}

func (s *GetBackupStorageCountResponseBodyBackupStorageCount) Validate() error {
	return dara.Validate(s)
}
