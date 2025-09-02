// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetMigrationProcessResponseBodyData) *GetMigrationProcessResponseBody
	GetData() []*GetMigrationProcessResponseBodyData
	SetErrorCode(v string) *GetMigrationProcessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMigrationProcessResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMigrationProcessResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMigrationProcessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMigrationProcessResponseBody
	GetSuccess() *bool
}

type GetMigrationProcessResponseBody struct {
	// The progress information of the migration task, including the names of all steps in and status of the migration task.
	Data []*GetMigrationProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 110001123456
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// test error msg
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// SADFSDFSD-SDFSDF-XDXCVX-ESWW
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMigrationProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationProcessResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationProcessResponseBody) GetData() []*GetMigrationProcessResponseBodyData {
	return s.Data
}

func (s *GetMigrationProcessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMigrationProcessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMigrationProcessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMigrationProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMigrationProcessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMigrationProcessResponseBody) SetData(v []*GetMigrationProcessResponseBodyData) *GetMigrationProcessResponseBody {
	s.Data = v
	return s
}

func (s *GetMigrationProcessResponseBody) SetErrorCode(v string) *GetMigrationProcessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMigrationProcessResponseBody) SetErrorMessage(v string) *GetMigrationProcessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMigrationProcessResponseBody) SetHttpStatusCode(v int32) *GetMigrationProcessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMigrationProcessResponseBody) SetRequestId(v string) *GetMigrationProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMigrationProcessResponseBody) SetSuccess(v bool) *GetMigrationProcessResponseBody {
	s.Success = &v
	return s
}

func (s *GetMigrationProcessResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMigrationProcessResponseBodyData struct {
	// The name of the step in the migration task.
	//
	// example:
	//
	// IMPORE_PREPARE
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The running status of the step in the migration task. Valid values:
	//
	// 	- INIT
	//
	// 	- RUNNING
	//
	// 	- FAILURE
	//
	// 	- SUCCESS
	//
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetMigrationProcessResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMigrationProcessResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *GetMigrationProcessResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetMigrationProcessResponseBodyData) SetTaskName(v string) *GetMigrationProcessResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *GetMigrationProcessResponseBodyData) SetTaskStatus(v string) *GetMigrationProcessResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetMigrationProcessResponseBodyData) Validate() error {
	return dara.Validate(s)
}
