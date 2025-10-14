// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRplInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateRplInspectionTaskResponseBodyData) *CreateRplInspectionTaskResponseBody
	GetData() *CreateRplInspectionTaskResponseBodyData
	SetMessage(v string) *CreateRplInspectionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRplInspectionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRplInspectionTaskResponseBody
	GetSuccess() *bool
}

type CreateRplInspectionTaskResponseBody struct {
	Data *CreateRplInspectionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRplInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRplInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRplInspectionTaskResponseBody) GetData() *CreateRplInspectionTaskResponseBodyData {
	return s.Data
}

func (s *CreateRplInspectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRplInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRplInspectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRplInspectionTaskResponseBody) SetData(v *CreateRplInspectionTaskResponseBodyData) *CreateRplInspectionTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateRplInspectionTaskResponseBody) SetMessage(v string) *CreateRplInspectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRplInspectionTaskResponseBody) SetRequestId(v string) *CreateRplInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRplInspectionTaskResponseBody) SetSuccess(v bool) *CreateRplInspectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRplInspectionTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRplInspectionTaskResponseBodyData struct {
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s CreateRplInspectionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateRplInspectionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRplInspectionTaskResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateRplInspectionTaskResponseBodyData) SetSlinkTaskId(v string) *CreateRplInspectionTaskResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateRplInspectionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
