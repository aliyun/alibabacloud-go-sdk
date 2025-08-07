// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
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
	// 11111111
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F3322AFE-083E-4D77-A074-421301******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *CreateBackupResponseBody) GetRequestId() *string {
	return s.RequestId
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
