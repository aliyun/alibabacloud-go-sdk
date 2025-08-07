// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRestoreTaskResponseBody
	GetRequestId() *string
	SetRestoreTaskId(v string) *CreateRestoreTaskResponseBody
	GetRestoreTaskId() *string
	SetStatus(v string) *CreateRestoreTaskResponseBody
	GetStatus() *string
}

type CreateRestoreTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the restore task.
	//
	// example:
	//
	// restore-fdsafda
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	// The status of the service instance. Valid values:
	//
	// 	- Restoring
	//
	// 	- Restored
	//
	// 	- RestoreFailed
	//
	// example:
	//
	// Restoring
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRestoreTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRestoreTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRestoreTaskResponseBody) GetRestoreTaskId() *string {
	return s.RestoreTaskId
}

func (s *CreateRestoreTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateRestoreTaskResponseBody) SetRequestId(v string) *CreateRestoreTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetRestoreTaskId(v string) *CreateRestoreTaskResponseBody {
	s.RestoreTaskId = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) SetStatus(v string) *CreateRestoreTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateRestoreTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
