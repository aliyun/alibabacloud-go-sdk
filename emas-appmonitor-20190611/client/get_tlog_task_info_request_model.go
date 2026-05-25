// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetTlogTaskInfoRequest
	GetAppKey() *int64
	SetOs(v string) *GetTlogTaskInfoRequest
	GetOs() *string
	SetTaskId(v string) *GetTlogTaskInfoRequest
	GetTaskId() *string
}

type GetTlogTaskInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 233588686
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WXAIGCOYIQELGKLAL6IFXECD4B
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTlogTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTlogTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTlogTaskInfoRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetTlogTaskInfoRequest) GetOs() *string {
	return s.Os
}

func (s *GetTlogTaskInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTlogTaskInfoRequest) SetAppKey(v int64) *GetTlogTaskInfoRequest {
	s.AppKey = &v
	return s
}

func (s *GetTlogTaskInfoRequest) SetOs(v string) *GetTlogTaskInfoRequest {
	s.Os = &v
	return s
}

func (s *GetTlogTaskInfoRequest) SetTaskId(v string) *GetTlogTaskInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetTlogTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
