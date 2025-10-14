// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyColumnarClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyColumnarClassResponseBodyData) *ModifyColumnarClassResponseBody
	GetData() *ModifyColumnarClassResponseBodyData
	SetRequestId(v string) *ModifyColumnarClassResponseBody
	GetRequestId() *string
}

type ModifyColumnarClassResponseBody struct {
	Data *ModifyColumnarClassResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-****-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyColumnarClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyColumnarClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyColumnarClassResponseBody) GetData() *ModifyColumnarClassResponseBodyData {
	return s.Data
}

func (s *ModifyColumnarClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyColumnarClassResponseBody) SetData(v *ModifyColumnarClassResponseBodyData) *ModifyColumnarClassResponseBody {
	s.Data = v
	return s
}

func (s *ModifyColumnarClassResponseBody) SetRequestId(v string) *ModifyColumnarClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyColumnarClassResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyColumnarClassResponseBodyData struct {
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyColumnarClassResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyColumnarClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyColumnarClassResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyColumnarClassResponseBodyData) SetTaskId(v int32) *ModifyColumnarClassResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyColumnarClassResponseBodyData) Validate() error {
	return dara.Validate(s)
}
