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
	// The backup task ID. For cluster instances created before December 5, 2023, the return value consists of the job IDs of individual child DB nodes. For example, for a cluster instance with two shards, the BackupJobID value is returned in the format of "10000,10001". For cluster instances created on or after December 5, 2023, all DB nodes in a cluster are described by a single job ID. You only need to use this single job ID.
	//
	// example:
	//
	// "1022,2480"
	BackupJobID *string `json:"BackupJobID,omitempty" xml:"BackupJobID,omitempty"`
	// The request ID.
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
