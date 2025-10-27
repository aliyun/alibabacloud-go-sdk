// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkflowResponseBody
	GetCode() *string
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that was returned for the request.
	Data *CreateWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// timetype is invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the workflow was created. Valid values:
	//
	// - true
	//
	// - false
	//
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

func (s *CreateWorkflowResponseBody) GetCode() *string {
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

func (s *CreateWorkflowResponseBody) SetCode(v string) *CreateWorkflowResponseBody {
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
	// The workflow ID.
	//
	// example:
	//
	// 92583
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
