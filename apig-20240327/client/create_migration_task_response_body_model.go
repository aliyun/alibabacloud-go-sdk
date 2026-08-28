// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMigrationTaskResponseBody
	GetCode() *string
	SetData(v *CreateMigrationTaskResponseBodyData) *CreateMigrationTaskResponseBody
	GetData() *CreateMigrationTaskResponseBodyData
	SetMessage(v string) *CreateMigrationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMigrationTaskResponseBody
	GetRequestId() *string
}

type CreateMigrationTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateMigrationTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 019FD4D8-8A86-5FDE-B79F-357C69677DFB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMigrationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMigrationTaskResponseBody) GetData() *CreateMigrationTaskResponseBodyData {
	return s.Data
}

func (s *CreateMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMigrationTaskResponseBody) SetCode(v string) *CreateMigrationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMigrationTaskResponseBody) SetData(v *CreateMigrationTaskResponseBodyData) *CreateMigrationTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateMigrationTaskResponseBody) SetMessage(v string) *CreateMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMigrationTaskResponseBody) SetRequestId(v string) *CreateMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMigrationTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMigrationTaskResponseBodyData struct {
	// example:
	//
	// async-task-xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateMigrationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMigrationTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMigrationTaskResponseBodyData) SetTaskId(v string) *CreateMigrationTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateMigrationTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
