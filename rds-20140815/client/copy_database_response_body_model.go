// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBName(v string) *CopyDatabaseResponseBody
	GetDBName() *string
	SetDBStatus(v string) *CopyDatabaseResponseBody
	GetDBStatus() *string
	SetRequestId(v string) *CopyDatabaseResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CopyDatabaseResponseBody
	GetTaskId() *string
}

type CopyDatabaseResponseBody struct {
	// The name of the database.
	//
	// example:
	//
	// test02
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The status of the database. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Running**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Creating
	DBStatus *string `json:"DBStatus,omitempty" xml:"DBStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 2562****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CopyDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDatabaseResponseBody) GetDBName() *string {
	return s.DBName
}

func (s *CopyDatabaseResponseBody) GetDBStatus() *string {
	return s.DBStatus
}

func (s *CopyDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyDatabaseResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CopyDatabaseResponseBody) SetDBName(v string) *CopyDatabaseResponseBody {
	s.DBName = &v
	return s
}

func (s *CopyDatabaseResponseBody) SetDBStatus(v string) *CopyDatabaseResponseBody {
	s.DBStatus = &v
	return s
}

func (s *CopyDatabaseResponseBody) SetRequestId(v string) *CopyDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyDatabaseResponseBody) SetTaskId(v string) *CopyDatabaseResponseBody {
	s.TaskId = &v
	return s
}

func (s *CopyDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
