// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseSkillPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ParseSkillPackageResponseBody
	GetRequestId() *string
	SetTaskKey(v string) *ParseSkillPackageResponseBody
	GetTaskKey() *string
}

type ParseSkillPackageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05C2791F-41A7-5E7C-B5E4-1401FD0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The key of the skill package parsing task.
	//
	// example:
	//
	// 2E7D8B71-2677-1B4C-9E25-A88B9******
	TaskKey *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
}

func (s ParseSkillPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ParseSkillPackageResponseBody) GoString() string {
	return s.String()
}

func (s *ParseSkillPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ParseSkillPackageResponseBody) GetTaskKey() *string {
	return s.TaskKey
}

func (s *ParseSkillPackageResponseBody) SetRequestId(v string) *ParseSkillPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ParseSkillPackageResponseBody) SetTaskKey(v string) *ParseSkillPackageResponseBody {
	s.TaskKey = &v
	return s
}

func (s *ParseSkillPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
