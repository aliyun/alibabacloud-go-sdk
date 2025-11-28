// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobId(v int64) *CreateBackupResponseBody
	GetBackupJobId() *int64
	SetRequestId(v string) *CreateBackupResponseBody
	GetRequestId() *string
}

type CreateBackupResponseBody struct {
	// The backup job ID.
	//
	// example:
	//
	// 123
	BackupJobId *int64 `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) GetBackupJobId() *int64 {
	return s.BackupJobId
}

func (s *CreateBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupResponseBody) SetBackupJobId(v int64) *CreateBackupResponseBody {
	s.BackupJobId = &v
	return s
}

func (s *CreateBackupResponseBody) SetRequestId(v string) *CreateBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
