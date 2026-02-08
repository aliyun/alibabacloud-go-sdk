// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMaterialTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskParam(v string) *SubmitMaterialTaskRequest
	GetTaskParam() *string
	SetTaskType(v string) *SubmitMaterialTaskRequest
	GetTaskType() *string
}

type SubmitMaterialTaskRequest struct {
	// This parameter is required.
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
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

func (s *SubmitMaterialTaskRequest) GetTaskParam() *string {
	return s.TaskParam
}

func (s *SubmitMaterialTaskRequest) GetTaskType() *string {
	return s.TaskType
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
