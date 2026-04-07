// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteMem0ResponseBodyData) *DeleteMem0ResponseBody
	GetData() *DeleteMem0ResponseBodyData
	SetRequestId(v string) *DeleteMem0ResponseBody
	GetRequestId() *string
}

type DeleteMem0ResponseBody struct {
	Data *DeleteMem0ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMem0ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMem0ResponseBody) GetData() *DeleteMem0ResponseBodyData {
	return s.Data
}

func (s *DeleteMem0ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMem0ResponseBody) SetData(v *DeleteMem0ResponseBodyData) *DeleteMem0ResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMem0ResponseBody) SetRequestId(v string) *DeleteMem0ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMem0ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMem0ResponseBodyData struct {
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteMem0ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0ResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteMem0ResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteMem0ResponseBodyData) SetTaskId(v int32) *DeleteMem0ResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteMem0ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
