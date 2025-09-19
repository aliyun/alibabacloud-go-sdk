// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupRestoreCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRestoreCount(v *DescribeBackupRestoreCountResponseBodyBackupRestoreCount) *DescribeBackupRestoreCountResponseBody
	GetBackupRestoreCount() *DescribeBackupRestoreCountResponseBodyBackupRestoreCount
	SetRequestId(v string) *DescribeBackupRestoreCountResponseBody
	GetRequestId() *string
}

type DescribeBackupRestoreCountResponseBody struct {
	// The statistics of restoration tasks.
	BackupRestoreCount *DescribeBackupRestoreCountResponseBodyBackupRestoreCount `json:"BackupRestoreCount,omitempty" xml:"BackupRestoreCount,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// ECC6B3E3-D496-512D-B46D-E6996A6B63EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupRestoreCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupRestoreCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupRestoreCountResponseBody) GetBackupRestoreCount() *DescribeBackupRestoreCountResponseBodyBackupRestoreCount {
	return s.BackupRestoreCount
}

func (s *DescribeBackupRestoreCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupRestoreCountResponseBody) SetBackupRestoreCount(v *DescribeBackupRestoreCountResponseBodyBackupRestoreCount) *DescribeBackupRestoreCountResponseBody {
	s.BackupRestoreCount = v
	return s
}

func (s *DescribeBackupRestoreCountResponseBody) SetRequestId(v string) *DescribeBackupRestoreCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupRestoreCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupRestoreCountResponseBodyBackupRestoreCount struct {
	// The number of the restoration tasks that are in the **being restored*	- state.
	//
	// example:
	//
	// 3
	Recovering *int32 `json:"Recovering,omitempty" xml:"Recovering,omitempty"`
	// The total number of the restoration tasks that you create.
	//
	// example:
	//
	// 30
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBackupRestoreCountResponseBodyBackupRestoreCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupRestoreCountResponseBodyBackupRestoreCount) GoString() string {
	return s.String()
}

func (s *DescribeBackupRestoreCountResponseBodyBackupRestoreCount) GetRecovering() *int32 {
	return s.Recovering
}

func (s *DescribeBackupRestoreCountResponseBodyBackupRestoreCount) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeBackupRestoreCountResponseBodyBackupRestoreCount) SetRecovering(v int32) *DescribeBackupRestoreCountResponseBodyBackupRestoreCount {
	s.Recovering = &v
	return s
}

func (s *DescribeBackupRestoreCountResponseBodyBackupRestoreCount) SetTotal(v int32) *DescribeBackupRestoreCountResponseBodyBackupRestoreCount {
	s.Total = &v
	return s
}

func (s *DescribeBackupRestoreCountResponseBodyBackupRestoreCount) Validate() error {
	return dara.Validate(s)
}
