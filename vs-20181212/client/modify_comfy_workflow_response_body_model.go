// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComfyWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyComfyWorkflowResponseBody
	GetCode() *string
	SetCreationTime(v string) *ModifyComfyWorkflowResponseBody
	GetCreationTime() *string
	SetDescription(v string) *ModifyComfyWorkflowResponseBody
	GetDescription() *string
	SetMessage(v string) *ModifyComfyWorkflowResponseBody
	GetMessage() *string
	SetName(v string) *ModifyComfyWorkflowResponseBody
	GetName() *string
	SetRequestId(v string) *ModifyComfyWorkflowResponseBody
	GetRequestId() *string
	SetUpdatedTime(v string) *ModifyComfyWorkflowResponseBody
	GetUpdatedTime() *string
	SetWorkflowId(v string) *ModifyComfyWorkflowResponseBody
	GetWorkflowId() *string
}

type ModifyComfyWorkflowResponseBody struct {
	// The error code. This parameter is returned only if the request fails.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-05-07T02:27:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The workflow description.
	//
	// example:
	//
	// 这是一个图生视频的工作流
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The returned message. This parameter provides error details if the request fails.
	//
	// example:
	//
	// conn failed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// 图生视频工作流示例
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-05-07T02:27:06Z
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// wf_adb32aed-ccdc-42ae-b4d4-a21181ac8a5f
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ModifyComfyWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyComfyWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyComfyWorkflowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyComfyWorkflowResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ModifyComfyWorkflowResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ModifyComfyWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyComfyWorkflowResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifyComfyWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyComfyWorkflowResponseBody) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ModifyComfyWorkflowResponseBody) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ModifyComfyWorkflowResponseBody) SetCode(v string) *ModifyComfyWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyComfyWorkflowResponseBody) SetCreationTime(v string) *ModifyComfyWorkflowResponseBody {
	s.CreationTime = &v
	return s
}

func (s *ModifyComfyWorkflowResponseBody) SetDescription(v string) *ModifyComfyWorkflowResponseBody {
	s.Description = &v
	return s
}

func (s *ModifyComfyWorkflowResponseBody) SetMessage(v string) *ModifyComfyWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyComfyWorkflowResponseBody) SetName(v string) *ModifyComfyWorkflowResponseBody {
	s.Name = &v
	return s
}

func (s *ModifyComfyWorkflowResponseBody) SetRequestId(v string) *ModifyComfyWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyComfyWorkflowResponseBody) SetUpdatedTime(v string) *ModifyComfyWorkflowResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *ModifyComfyWorkflowResponseBody) SetWorkflowId(v string) *ModifyComfyWorkflowResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *ModifyComfyWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
