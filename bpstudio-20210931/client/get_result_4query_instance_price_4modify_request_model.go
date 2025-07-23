// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResult4QueryInstancePrice4ModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetResult4QueryInstancePrice4ModifyRequest
	GetApplicationId() *string
	SetTaskId(v string) *GetResult4QueryInstancePrice4ModifyRequest
	GetTaskId() *string
}

type GetResult4QueryInstancePrice4ModifyRequest struct {
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 58d5d8c5b5489150417a7cd6caa614bb
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyRequest) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetResult4QueryInstancePrice4ModifyRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetResult4QueryInstancePrice4ModifyRequest) SetApplicationId(v string) *GetResult4QueryInstancePrice4ModifyRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyRequest) SetTaskId(v string) *GetResult4QueryInstancePrice4ModifyRequest {
	s.TaskId = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyRequest) Validate() error {
	return dara.Validate(s)
}
