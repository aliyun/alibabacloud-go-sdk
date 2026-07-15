// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFreeSize(v int64) *DescribeBackupStorageResponseBody
	GetFreeSize() *int64
	SetFullStorageSize(v int64) *DescribeBackupStorageResponseBody
	GetFullStorageSize() *int64
	SetLogStorageSize(v int64) *DescribeBackupStorageResponseBody
	GetLogStorageSize() *int64
	SetRequestId(v string) *DescribeBackupStorageResponseBody
	GetRequestId() *string
}

type DescribeBackupStorageResponseBody struct {
	// The free backup quota for the instance. Unit: bytes.
	//
	// example:
	//
	// 42949672960
	FreeSize *int64 `json:"FreeSize,omitempty" xml:"FreeSize,omitempty"`
	// The storage space used by full backups. Unit: bytes.
	//
	// > Instances that use cloud disks are backed up using snapshots. The size of a full backup is the total size of the snapshot chain.
	//
	// example:
	//
	// 789221621
	FullStorageSize *int64 `json:"FullStorageSize,omitempty" xml:"FullStorageSize,omitempty"`
	// The storage space used by log backups. Unit: bytes.
	//
	// example:
	//
	// 7892216
	LogStorageSize *int64 `json:"LogStorageSize,omitempty" xml:"LogStorageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D648B367-15B6-1B42-BD4B-4XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupStorageResponseBody) GetFreeSize() *int64 {
	return s.FreeSize
}

func (s *DescribeBackupStorageResponseBody) GetFullStorageSize() *int64 {
	return s.FullStorageSize
}

func (s *DescribeBackupStorageResponseBody) GetLogStorageSize() *int64 {
	return s.LogStorageSize
}

func (s *DescribeBackupStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupStorageResponseBody) SetFreeSize(v int64) *DescribeBackupStorageResponseBody {
	s.FreeSize = &v
	return s
}

func (s *DescribeBackupStorageResponseBody) SetFullStorageSize(v int64) *DescribeBackupStorageResponseBody {
	s.FullStorageSize = &v
	return s
}

func (s *DescribeBackupStorageResponseBody) SetLogStorageSize(v int64) *DescribeBackupStorageResponseBody {
	s.LogStorageSize = &v
	return s
}

func (s *DescribeBackupStorageResponseBody) SetRequestId(v string) *DescribeBackupStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
