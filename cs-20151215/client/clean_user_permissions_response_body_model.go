// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanUserPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CleanUserPermissionsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CleanUserPermissionsResponseBody
	GetTaskId() *string
}

type CleanUserPermissionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 687C5BAA-D103-4993-884B-C35E4314****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// clean-user-permissions-2085266204********-6694c16e6ae07***********
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CleanUserPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CleanUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *CleanUserPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CleanUserPermissionsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CleanUserPermissionsResponseBody) SetRequestId(v string) *CleanUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CleanUserPermissionsResponseBody) SetTaskId(v string) *CleanUserPermissionsResponseBody {
	s.TaskId = &v
	return s
}

func (s *CleanUserPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
