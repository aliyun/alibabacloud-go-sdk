// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCNInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateSubCNInstanceResponseBodyData) *CreateSubCNInstanceResponseBody
	GetData() *CreateSubCNInstanceResponseBodyData
	SetRequestId(v string) *CreateSubCNInstanceResponseBody
	GetRequestId() *string
}

type CreateSubCNInstanceResponseBody struct {
	Data *CreateSubCNInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// D6A4256F-7B83-5BD7-9AC0-72E1FAC05330
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSubCNInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCNInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubCNInstanceResponseBody) GetData() *CreateSubCNInstanceResponseBodyData {
	return s.Data
}

func (s *CreateSubCNInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubCNInstanceResponseBody) SetData(v *CreateSubCNInstanceResponseBodyData) *CreateSubCNInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateSubCNInstanceResponseBody) SetRequestId(v string) *CreateSubCNInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubCNInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSubCNInstanceResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSubCNInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCNInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSubCNInstanceResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateSubCNInstanceResponseBodyData) SetTaskId(v int32) *CreateSubCNInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateSubCNInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
