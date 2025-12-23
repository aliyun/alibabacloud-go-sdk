// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateWorkflowResponseBodyData) *CreateWorkflowResponseBody
	GetData() *CreateWorkflowResponseBodyData
	SetRequestId(v string) *CreateWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkflowResponseBody
	GetSuccess() *bool
}

type CreateWorkflowResponseBody struct {
	Data *CreateWorkflowResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA38***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponseBody) GetData() *CreateWorkflowResponseBodyData {
	return s.Data
}

func (s *CreateWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkflowResponseBody) SetData(v *CreateWorkflowResponseBodyData) *CreateWorkflowResponseBody {
	s.Data = v
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
	// w-acfmv4opbs****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
}

func (s CreateWorkflowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkflowResponseBodyData) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *CreateWorkflowResponseBodyData) SetWorkflowId(v string) *CreateWorkflowResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *CreateWorkflowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
