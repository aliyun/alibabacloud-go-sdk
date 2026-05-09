// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubCNInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteSubCNInstanceResponseBodyData) *DeleteSubCNInstanceResponseBody
	GetData() *DeleteSubCNInstanceResponseBodyData
	SetRequestId(v string) *DeleteSubCNInstanceResponseBody
	GetRequestId() *string
}

type DeleteSubCNInstanceResponseBody struct {
	Data *DeleteSubCNInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1A586DCB-39A6-4050-81CC-C7BD4CCDB49F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubCNInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubCNInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubCNInstanceResponseBody) GetData() *DeleteSubCNInstanceResponseBodyData {
	return s.Data
}

func (s *DeleteSubCNInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubCNInstanceResponseBody) SetData(v *DeleteSubCNInstanceResponseBodyData) *DeleteSubCNInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSubCNInstanceResponseBody) SetRequestId(v string) *DeleteSubCNInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubCNInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSubCNInstanceResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteSubCNInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubCNInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSubCNInstanceResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteSubCNInstanceResponseBodyData) SetTaskId(v int32) *DeleteSubCNInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteSubCNInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
