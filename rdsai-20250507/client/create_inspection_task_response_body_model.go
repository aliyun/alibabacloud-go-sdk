// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateInspectionTaskResponseBodyData) *CreateInspectionTaskResponseBody
	GetData() *CreateInspectionTaskResponseBodyData
	SetMessage(v string) *CreateInspectionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInspectionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInspectionTaskResponseBody
	GetSuccess() *bool
}

type CreateInspectionTaskResponseBody struct {
	Data    *CreateInspectionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInspectionTaskResponseBody) GetData() *CreateInspectionTaskResponseBodyData {
	return s.Data
}

func (s *CreateInspectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInspectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInspectionTaskResponseBody) SetData(v *CreateInspectionTaskResponseBodyData) *CreateInspectionTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateInspectionTaskResponseBody) SetMessage(v string) *CreateInspectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInspectionTaskResponseBody) SetRequestId(v string) *CreateInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInspectionTaskResponseBody) SetSuccess(v bool) *CreateInspectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInspectionTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInspectionTaskResponseBodyData struct {
	// example:
	//
	// 9adf8567-b619-4d37-8ff2-01d38a76****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateInspectionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInspectionTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateInspectionTaskResponseBodyData) SetTaskId(v string) *CreateInspectionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateInspectionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
