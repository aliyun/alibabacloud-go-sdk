// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRecordTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopRecordTaskRequest
	GetAppId() *string
	SetOwnerId(v int64) *StopRecordTaskRequest
	GetOwnerId() *int64
	SetTaskId(v string) *StopRecordTaskRequest
	GetTaskId() *string
}

type StopRecordTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRecordTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *StopRecordTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopRecordTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopRecordTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopRecordTaskRequest) SetAppId(v string) *StopRecordTaskRequest {
	s.AppId = &v
	return s
}

func (s *StopRecordTaskRequest) SetOwnerId(v int64) *StopRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopRecordTaskRequest) SetTaskId(v string) *StopRecordTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopRecordTaskRequest) Validate() error {
	return dara.Validate(s)
}
