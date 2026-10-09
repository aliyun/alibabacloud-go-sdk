// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobId(v int64) *CreateSupabaseBackupResponseBody
	GetBackupJobId() *int64
	SetRequestId(v string) *CreateSupabaseBackupResponseBody
	GetRequestId() *string
}

type CreateSupabaseBackupResponseBody struct {
	// The ID of the backup job. You can call ListSupabaseBackupJobs to query the status and progress of the corresponding job.
	//
	// example:
	//
	// 123
	BackupJobId *int64 `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSupabaseBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSupabaseBackupResponseBody) GetBackupJobId() *int64 {
	return s.BackupJobId
}

func (s *CreateSupabaseBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSupabaseBackupResponseBody) SetBackupJobId(v int64) *CreateSupabaseBackupResponseBody {
	s.BackupJobId = &v
	return s
}

func (s *CreateSupabaseBackupResponseBody) SetRequestId(v string) *CreateSupabaseBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSupabaseBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
