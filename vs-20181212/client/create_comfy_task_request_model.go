// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *CreateComfyTaskRequest
	GetHiveId() *string
	SetUserParameters(v string) *CreateComfyTaskRequest
	GetUserParameters() *string
	SetWorkflowId(v string) *CreateComfyTaskRequest
	GetWorkflowId() *string
}

type CreateComfyTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hive-26cd567b35c04a0a90f017388207b232
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"2":{"text":"masterpiece, best quality, beautiful cat, sunny day"},"3":{"text":"low quality, worst quality, blurry, watermark, text, signature"}}
	UserParameters *string `json:"UserParameters,omitempty" xml:"UserParameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// wf_adb32aed-ccdc-42ae-b4d4-a21181ac8a5f
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s CreateComfyTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateComfyTaskRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *CreateComfyTaskRequest) GetUserParameters() *string {
	return s.UserParameters
}

func (s *CreateComfyTaskRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *CreateComfyTaskRequest) SetHiveId(v string) *CreateComfyTaskRequest {
	s.HiveId = &v
	return s
}

func (s *CreateComfyTaskRequest) SetUserParameters(v string) *CreateComfyTaskRequest {
	s.UserParameters = &v
	return s
}

func (s *CreateComfyTaskRequest) SetWorkflowId(v string) *CreateComfyTaskRequest {
	s.WorkflowId = &v
	return s
}

func (s *CreateComfyTaskRequest) Validate() error {
	return dara.Validate(s)
}
