// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *CreateAICoachTaskSessionRequest
	GetTaskId() *string
	SetUid(v string) *CreateAICoachTaskSessionRequest
	GetUid() *string
}

type CreateAICoachTaskSessionRequest struct {
	// example:
	//
	// 821882330423951360
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 1730530943640489
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateAICoachTaskSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskSessionRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAICoachTaskSessionRequest) GetUid() *string {
	return s.Uid
}

func (s *CreateAICoachTaskSessionRequest) SetTaskId(v string) *CreateAICoachTaskSessionRequest {
	s.TaskId = &v
	return s
}

func (s *CreateAICoachTaskSessionRequest) SetUid(v string) *CreateAICoachTaskSessionRequest {
	s.Uid = &v
	return s
}

func (s *CreateAICoachTaskSessionRequest) Validate() error {
	return dara.Validate(s)
}
