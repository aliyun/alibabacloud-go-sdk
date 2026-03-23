// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *GetBatchTaskProgressRequest
	GetAppCode() *string
	SetAppName(v string) *GetBatchTaskProgressRequest
	GetAppName() *string
	SetTaskId(v string) *GetBatchTaskProgressRequest
	GetTaskId() *string
}

type GetBatchTaskProgressRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetBatchTaskProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTaskProgressRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetBatchTaskProgressRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetBatchTaskProgressRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetBatchTaskProgressRequest) SetAppCode(v string) *GetBatchTaskProgressRequest {
	s.AppCode = &v
	return s
}

func (s *GetBatchTaskProgressRequest) SetAppName(v string) *GetBatchTaskProgressRequest {
	s.AppName = &v
	return s
}

func (s *GetBatchTaskProgressRequest) SetTaskId(v string) *GetBatchTaskProgressRequest {
	s.TaskId = &v
	return s
}

func (s *GetBatchTaskProgressRequest) Validate() error {
	return dara.Validate(s)
}
