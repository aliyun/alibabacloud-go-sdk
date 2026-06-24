// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelTaskRequest
	GetClientToken() *string
	SetTaskType(v string) *CancelTaskRequest
	GetTaskType() *string
}

type CancelTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The task type. Set this parameter to MigrateData, which indicates a data migration task.
	//
	// This parameter is required.
	//
	// example:
	//
	// MigrateData
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CancelTaskRequest) SetClientToken(v string) *CancelTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelTaskRequest) SetTaskType(v string) *CancelTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CancelTaskRequest) Validate() error {
	return dara.Validate(s)
}
