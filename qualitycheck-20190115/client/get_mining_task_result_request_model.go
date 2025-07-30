// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiningTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetMiningTaskResultRequest
	GetBaseMeAgentId() *int64
	SetTaskId(v string) *GetMiningTaskResultRequest
	GetTaskId() *string
}

type GetMiningTaskResultRequest struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetMiningTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMiningTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetMiningTaskResultRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetMiningTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMiningTaskResultRequest) SetBaseMeAgentId(v int64) *GetMiningTaskResultRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetMiningTaskResultRequest) SetTaskId(v string) *GetMiningTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetMiningTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
