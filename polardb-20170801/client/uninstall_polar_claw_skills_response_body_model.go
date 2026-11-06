// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPolarClawSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallPolarClawSkillsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UninstallPolarClawSkillsResponseBody
	GetTaskId() *string
}

type UninstallPolarClawSkillsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 123456
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UninstallPolarClawSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallPolarClawSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallPolarClawSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallPolarClawSkillsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallPolarClawSkillsResponseBody) SetRequestId(v string) *UninstallPolarClawSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallPolarClawSkillsResponseBody) SetTaskId(v string) *UninstallPolarClawSkillsResponseBody {
	s.TaskId = &v
	return s
}

func (s *UninstallPolarClawSkillsResponseBody) Validate() error {
	return dara.Validate(s)
}
