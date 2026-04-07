// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMem0ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateMem0ResponseBodyData) *CreateMem0ResponseBody
	GetData() *CreateMem0ResponseBodyData
	SetRequestId(v string) *CreateMem0ResponseBody
	GetRequestId() *string
}

type CreateMem0ResponseBody struct {
	Data *CreateMem0ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMem0ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMem0ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMem0ResponseBody) GetData() *CreateMem0ResponseBodyData {
	return s.Data
}

func (s *CreateMem0ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMem0ResponseBody) SetData(v *CreateMem0ResponseBodyData) *CreateMem0ResponseBody {
	s.Data = v
	return s
}

func (s *CreateMem0ResponseBody) SetRequestId(v string) *CreateMem0ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMem0ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMem0ResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateMem0ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMem0ResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMem0ResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateMem0ResponseBodyData) SetTaskId(v int32) *CreateMem0ResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateMem0ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
