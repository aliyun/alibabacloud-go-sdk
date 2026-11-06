// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobId(v string) *CreateApplicationBackupResponseBody
	GetBackupJobId() *string
	SetRequestId(v string) *CreateApplicationBackupResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateApplicationBackupResponseBody
	GetTaskId() *string
}

type CreateApplicationBackupResponseBody struct {
	// The ID of the created backup job.
	//
	// example:
	//
	// 11111111
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the workflow task corresponding to the backup creation.
	//
	// example:
	//
	// 32077061
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateApplicationBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationBackupResponseBody) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *CreateApplicationBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationBackupResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateApplicationBackupResponseBody) SetBackupJobId(v string) *CreateApplicationBackupResponseBody {
	s.BackupJobId = &v
	return s
}

func (s *CreateApplicationBackupResponseBody) SetRequestId(v string) *CreateApplicationBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationBackupResponseBody) SetTaskId(v string) *CreateApplicationBackupResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateApplicationBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
