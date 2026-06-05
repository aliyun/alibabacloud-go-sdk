// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateComfyWorkflowResponseBody
	GetCode() *int64
	SetMessage(v string) *CreateComfyWorkflowResponseBody
	GetMessage() *string
	SetMissingNodes(v []*string) *CreateComfyWorkflowResponseBody
	GetMissingNodes() []*string
	SetRequestId(v string) *CreateComfyWorkflowResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateComfyWorkflowResponseBody
	GetStatus() *string
	SetWorkflowId(v string) *CreateComfyWorkflowResponseBody
	GetWorkflowId() *string
}

type CreateComfyWorkflowResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// conn failed!
	Message      *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	MissingNodes []*string `json:"MissingNodes,omitempty" xml:"MissingNodes,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// wf_adb32aed-ccdc-42ae-b4d4-a21181ac8a5c
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s CreateComfyWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComfyWorkflowResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateComfyWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateComfyWorkflowResponseBody) GetMissingNodes() []*string {
	return s.MissingNodes
}

func (s *CreateComfyWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateComfyWorkflowResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateComfyWorkflowResponseBody) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *CreateComfyWorkflowResponseBody) SetCode(v int64) *CreateComfyWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComfyWorkflowResponseBody) SetMessage(v string) *CreateComfyWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateComfyWorkflowResponseBody) SetMissingNodes(v []*string) *CreateComfyWorkflowResponseBody {
	s.MissingNodes = v
	return s
}

func (s *CreateComfyWorkflowResponseBody) SetRequestId(v string) *CreateComfyWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateComfyWorkflowResponseBody) SetStatus(v string) *CreateComfyWorkflowResponseBody {
	s.Status = &v
	return s
}

func (s *CreateComfyWorkflowResponseBody) SetWorkflowId(v string) *CreateComfyWorkflowResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *CreateComfyWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
