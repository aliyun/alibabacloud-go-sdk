// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstanceResponseBody
	GetCode() *string
	SetData(v *DeleteInstanceResponseBodyData) *DeleteInstanceResponseBody
	GetData() *DeleteInstanceResponseBodyData
	SetMessage(v string) *DeleteInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstanceResponseBody
	GetSuccess() *bool
}

type DeleteInstanceResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// req-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceResponseBody) GetData() *DeleteInstanceResponseBodyData {
	return s.Data
}

func (s *DeleteInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceResponseBody) SetCode(v string) *DeleteInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetData(v *DeleteInstanceResponseBodyData) *DeleteInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteInstanceResponseBody) SetMessage(v string) *DeleteInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteInstanceResponseBodyData struct {
	// example:
	//
	// inst-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// task-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteInstanceResponseBodyData) SetInstanceId(v string) *DeleteInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceResponseBodyData) SetTaskId(v string) *DeleteInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
