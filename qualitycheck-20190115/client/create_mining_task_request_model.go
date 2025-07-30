// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMiningTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateMiningTaskRequest
	GetBaseMeAgentId() *int64
	SetCallbackUrl(v string) *CreateMiningTaskRequest
	GetCallbackUrl() *string
	SetFilePath(v string) *CreateMiningTaskRequest
	GetFilePath() *string
	SetParam(v string) *CreateMiningTaskRequest
	GetParam() *string
	SetTaskType(v string) *CreateMiningTaskRequest
	GetTaskType() *string
}

type CreateMiningTaskRequest struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64  `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	CallbackUrl   *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// 123.22.com/11.csv
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// {"startDate":"20250505"}
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// example:
	//
	// demandMining
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateMiningTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMiningTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMiningTaskRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateMiningTaskRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateMiningTaskRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateMiningTaskRequest) GetParam() *string {
	return s.Param
}

func (s *CreateMiningTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateMiningTaskRequest) SetBaseMeAgentId(v int64) *CreateMiningTaskRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateMiningTaskRequest) SetCallbackUrl(v string) *CreateMiningTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateMiningTaskRequest) SetFilePath(v string) *CreateMiningTaskRequest {
	s.FilePath = &v
	return s
}

func (s *CreateMiningTaskRequest) SetParam(v string) *CreateMiningTaskRequest {
	s.Param = &v
	return s
}

func (s *CreateMiningTaskRequest) SetTaskType(v string) *CreateMiningTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateMiningTaskRequest) Validate() error {
	return dara.Validate(s)
}
