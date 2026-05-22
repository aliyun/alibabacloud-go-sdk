// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUploadTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListUploadTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*ListUploadTasksResponseBodyTasks) *ListUploadTasksResponseBody
	GetTasks() []*ListUploadTasksResponseBodyTasks
}

type ListUploadTasksResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ListUploadTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListUploadTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUploadTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListUploadTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUploadTasksResponseBody) GetTasks() []*ListUploadTasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListUploadTasksResponseBody) SetRequestId(v string) *ListUploadTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUploadTasksResponseBody) SetTasks(v []*ListUploadTasksResponseBodyTasks) *ListUploadTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListUploadTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUploadTasksResponseBodyTasks struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UploadId       *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	UploadTaskName *string `json:"UploadTaskName,omitempty" xml:"UploadTaskName,omitempty"`
}

func (s ListUploadTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListUploadTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListUploadTasksResponseBodyTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUploadTasksResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *ListUploadTasksResponseBodyTasks) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUploadTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListUploadTasksResponseBodyTasks) GetType() *string {
	return s.Type
}

func (s *ListUploadTasksResponseBodyTasks) GetUploadId() *string {
	return s.UploadId
}

func (s *ListUploadTasksResponseBodyTasks) GetUploadTaskName() *string {
	return s.UploadTaskName
}

func (s *ListUploadTasksResponseBodyTasks) SetCreateTime(v string) *ListUploadTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetDescription(v string) *ListUploadTasksResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetErrorCode(v string) *ListUploadTasksResponseBodyTasks {
	s.ErrorCode = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetStatus(v string) *ListUploadTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetType(v string) *ListUploadTasksResponseBodyTasks {
	s.Type = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetUploadId(v string) *ListUploadTasksResponseBodyTasks {
	s.UploadId = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) SetUploadTaskName(v string) *ListUploadTasksResponseBodyTasks {
	s.UploadTaskName = &v
	return s
}

func (s *ListUploadTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
