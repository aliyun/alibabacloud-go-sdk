// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillFileDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHashKey(v string) *CreateSkillFileDetectResponseBody
	GetHashKey() *string
	SetRequestId(v string) *CreateSkillFileDetectResponseBody
	GetRequestId() *string
}

type CreateSkillFileDetectResponseBody struct {
	// The unique identifier for the detection task.
	//
	// example:
	//
	// 2aceb074-fa72-44d2-99d9-45b17cffe0e7
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F976EF8-C10A-57DC-917C-BB7BEB508FFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSkillFileDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileDetectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillFileDetectResponseBody) GetHashKey() *string {
	return s.HashKey
}

func (s *CreateSkillFileDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillFileDetectResponseBody) SetHashKey(v string) *CreateSkillFileDetectResponseBody {
	s.HashKey = &v
	return s
}

func (s *CreateSkillFileDetectResponseBody) SetRequestId(v string) *CreateSkillFileDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillFileDetectResponseBody) Validate() error {
	return dara.Validate(s)
}
