// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePolarClawSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradePolarClawSkillsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpgradePolarClawSkillsResponseBody
	GetTaskId() *string
}

type UpgradePolarClawSkillsResponseBody struct {
	// example:
	//
	// 45D24263-7E3A-4140-9472-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 32077061
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradePolarClawSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradePolarClawSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradePolarClawSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradePolarClawSkillsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradePolarClawSkillsResponseBody) SetRequestId(v string) *UpgradePolarClawSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradePolarClawSkillsResponseBody) SetTaskId(v string) *UpgradePolarClawSkillsResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradePolarClawSkillsResponseBody) Validate() error {
	return dara.Validate(s)
}
