// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobID(v string) *CreateBackupResponseBody
	GetBackupJobID() *string
	SetRequestId(v string) *CreateBackupResponseBody
	GetRequestId() *string
}

type CreateBackupResponseBody struct {
	// The ID of the backup task.
	//
	// example:
	//
	// 1162****
	BackupJobID *string `json:"BackupJobID,omitempty" xml:"BackupJobID,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2FF6158E-3394-4A90-B634-79C49184****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) GetBackupJobID() *string {
	return s.BackupJobID
}

func (s *CreateBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupResponseBody) SetBackupJobID(v string) *CreateBackupResponseBody {
	s.BackupJobID = &v
	return s
}

func (s *CreateBackupResponseBody) SetRequestId(v string) *CreateBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
