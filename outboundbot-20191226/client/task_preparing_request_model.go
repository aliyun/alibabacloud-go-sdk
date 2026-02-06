// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskPreparingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *TaskPreparingRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *TaskPreparingRequest
	GetInstanceOwnerId() *int64
	SetJobId(v string) *TaskPreparingRequest
	GetJobId() *string
}

type TaskPreparingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 33040b9a-b04b-452f-b554-cd6f3a15f850
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1971226538081821
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c9e42cd7-ba99-4872-9802-e05719ab051c
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s TaskPreparingRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskPreparingRequest) GoString() string {
	return s.String()
}

func (s *TaskPreparingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TaskPreparingRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *TaskPreparingRequest) GetJobId() *string {
	return s.JobId
}

func (s *TaskPreparingRequest) SetInstanceId(v string) *TaskPreparingRequest {
	s.InstanceId = &v
	return s
}

func (s *TaskPreparingRequest) SetInstanceOwnerId(v int64) *TaskPreparingRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *TaskPreparingRequest) SetJobId(v string) *TaskPreparingRequest {
	s.JobId = &v
	return s
}

func (s *TaskPreparingRequest) Validate() error {
	return dara.Validate(s)
}
