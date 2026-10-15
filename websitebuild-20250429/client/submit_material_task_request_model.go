// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMaterialTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizGroupId(v string) *SubmitMaterialTaskRequest
	GetBizGroupId() *string
	SetTaskParam(v string) *SubmitMaterialTaskRequest
	GetTaskParam() *string
	SetTaskType(v string) *SubmitMaterialTaskRequest
	GetTaskType() *string
}

type SubmitMaterialTaskRequest struct {
	// The business group ID.
	//
	// example:
	//
	// xxx
	BizGroupId *string `json:"BizGroupId,omitempty" xml:"BizGroupId,omitempty"`
	// The task parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "brandName": "Brand name",
	//
	//     "direction": "HORIZONTAL",
	//
	//     "userPrompt": "Design requirements"
	//
	// }
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	// The task type.
	//
	// This parameter is required.
	//
	// example:
	//
	// IMAGE_LOGO
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s SubmitMaterialTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMaterialTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitMaterialTaskRequest) GetBizGroupId() *string {
	return s.BizGroupId
}

func (s *SubmitMaterialTaskRequest) GetTaskParam() *string {
	return s.TaskParam
}

func (s *SubmitMaterialTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *SubmitMaterialTaskRequest) SetBizGroupId(v string) *SubmitMaterialTaskRequest {
	s.BizGroupId = &v
	return s
}

func (s *SubmitMaterialTaskRequest) SetTaskParam(v string) *SubmitMaterialTaskRequest {
	s.TaskParam = &v
	return s
}

func (s *SubmitMaterialTaskRequest) SetTaskType(v string) *SubmitMaterialTaskRequest {
	s.TaskType = &v
	return s
}

func (s *SubmitMaterialTaskRequest) Validate() error {
	return dara.Validate(s)
}
