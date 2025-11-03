// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *CreateBackupResponseBody
	GetBackupId() *string
	SetBackupJobId(v string) *CreateBackupResponseBody
	GetBackupJobId() *string
	SetRequestId(v string) *CreateBackupResponseBody
	GetRequestId() *string
}

type CreateBackupResponseBody struct {
	// The ID of the backup set.
	//
	// example:
	//
	// 5664****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// 775051
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7016B12F-7F64-40A4-BAFF-013F02AC82FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateBackupResponseBody) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *CreateBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupResponseBody) SetBackupId(v string) *CreateBackupResponseBody {
	s.BackupId = &v
	return s
}

func (s *CreateBackupResponseBody) SetBackupJobId(v string) *CreateBackupResponseBody {
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
