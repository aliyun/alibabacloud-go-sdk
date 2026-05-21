// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *CreateTempFileRequest
	GetCapacity() *int64
	SetInstanceId(v string) *CreateTempFileRequest
	GetInstanceId() *string
	SetName(v string) *CreateTempFileRequest
	GetName() *string
	SetPrefix(v string) *CreateTempFileRequest
	GetPrefix() *string
	SetTaskId(v string) *CreateTempFileRequest
	GetTaskId() *string
}

type CreateTempFileRequest struct {
	// example:
	//
	// 1000
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// filename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// somedir/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// task-05cexxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateTempFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileRequest) GoString() string {
	return s.String()
}

func (s *CreateTempFileRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CreateTempFileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTempFileRequest) GetName() *string {
	return s.Name
}

func (s *CreateTempFileRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *CreateTempFileRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTempFileRequest) SetCapacity(v int64) *CreateTempFileRequest {
	s.Capacity = &v
	return s
}

func (s *CreateTempFileRequest) SetInstanceId(v string) *CreateTempFileRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTempFileRequest) SetName(v string) *CreateTempFileRequest {
	s.Name = &v
	return s
}

func (s *CreateTempFileRequest) SetPrefix(v string) *CreateTempFileRequest {
	s.Prefix = &v
	return s
}

func (s *CreateTempFileRequest) SetTaskId(v string) *CreateTempFileRequest {
	s.TaskId = &v
	return s
}

func (s *CreateTempFileRequest) Validate() error {
	return dara.Validate(s)
}
