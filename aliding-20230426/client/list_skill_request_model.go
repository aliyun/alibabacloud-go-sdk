// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListSkillRequest
	GetGroupId() *string
}

type ListSkillRequest struct {
	// example:
	//
	// qweq-1231-jbarr-9940-asdf
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s ListSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillRequest) GoString() string {
	return s.String()
}

func (s *ListSkillRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListSkillRequest) SetGroupId(v string) *ListSkillRequest {
	s.GroupId = &v
	return s
}

func (s *ListSkillRequest) Validate() error {
	return dara.Validate(s)
}
