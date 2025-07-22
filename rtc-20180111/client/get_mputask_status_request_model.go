// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMPUTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMPUTaskStatusRequest
	GetAppId() *string
	SetOwnerId(v int64) *GetMPUTaskStatusRequest
	GetOwnerId() *int64
	SetTaskId(v string) *GetMPUTaskStatusRequest
	GetTaskId() *string
}

type GetMPUTaskStatusRequest struct {
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

func (s GetMPUTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMPUTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMPUTaskStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetMPUTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMPUTaskStatusRequest) SetAppId(v string) *GetMPUTaskStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetMPUTaskStatusRequest) SetOwnerId(v int64) *GetMPUTaskStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMPUTaskStatusRequest) SetTaskId(v string) *GetMPUTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetMPUTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
