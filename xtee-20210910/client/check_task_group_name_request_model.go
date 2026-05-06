// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTaskGroupNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckTaskGroupNameRequest
	GetLang() *string
	SetRegId(v string) *CheckTaskGroupNameRequest
	GetRegId() *string
	SetTaskGroupName(v string) *CheckTaskGroupNameRequest
	GetTaskGroupName() *string
}

type CheckTaskGroupNameRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// example:
	//
	// GroupNameTest
	TaskGroupName *string `json:"TaskGroupName,omitempty" xml:"TaskGroupName,omitempty"`
}

func (s CheckTaskGroupNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckTaskGroupNameRequest) GoString() string {
	return s.String()
}

func (s *CheckTaskGroupNameRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckTaskGroupNameRequest) GetRegId() *string {
	return s.RegId
}

func (s *CheckTaskGroupNameRequest) GetTaskGroupName() *string {
	return s.TaskGroupName
}

func (s *CheckTaskGroupNameRequest) SetLang(v string) *CheckTaskGroupNameRequest {
	s.Lang = &v
	return s
}

func (s *CheckTaskGroupNameRequest) SetRegId(v string) *CheckTaskGroupNameRequest {
	s.RegId = &v
	return s
}

func (s *CheckTaskGroupNameRequest) SetTaskGroupName(v string) *CheckTaskGroupNameRequest {
	s.TaskGroupName = &v
	return s
}

func (s *CheckTaskGroupNameRequest) Validate() error {
	return dara.Validate(s)
}
