// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOtaTaskByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetOtaTaskByTaskIdRequest
	GetTaskId() *string
}

type GetOtaTaskByTaskIdRequest struct {
	// The ID of the OTA update task. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ota-be7jzm29wrrz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetOtaTaskByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOtaTaskByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOtaTaskByTaskIdRequest) SetTaskId(v string) *GetOtaTaskByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *GetOtaTaskByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
