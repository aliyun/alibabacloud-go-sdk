// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateWorkflowResponseBody
	GetCode() *int32
	SetData(v *CreateWorkflowResponseBodyData) *CreateWorkflowResponseBody
	GetData() *CreateWorkflowResponseBodyData
	SetMessage(v string) *CreateWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkflowResponseBody
	GetSuccess() *bool
}

type CreateWorkflowResponseBody struct {
	// example:
	//
	// 200
	Code *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateWorkflowResponseBody) GetData() *CreateWorkflowResponseBodyData {
	return s.Data
}

func (s *CreateWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkflowResponseBody) SetCode(v int32) *CreateWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkflowResponseBody) SetData(v *CreateWorkflowResponseBodyData) *CreateWorkflowResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkflowResponseBody) SetMessage(v string) *CreateWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkflowResponseBody) SetRequestId(v string) *CreateWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkflowResponseBody) SetSuccess(v bool) *CreateWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkflowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkflowResponseBodyData struct {
	// example:
	//
	// 10
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s CreateWorkflowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponseBodyData) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *CreateWorkflowResponseBodyData) SetWorkflowId(v int64) *CreateWorkflowResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *CreateWorkflowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
