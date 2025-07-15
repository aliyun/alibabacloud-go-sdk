// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingChatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListWaitingChatsRequest
	GetInstanceId() *string
	SetSkillGroupIdList(v string) *ListWaitingChatsRequest
	GetSkillGroupIdList() *string
}

type ListWaitingChatsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["skillgroup1@ccc-test", "skillgroup2@ccc-test"]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s ListWaitingChatsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingChatsRequest) GoString() string {
	return s.String()
}

func (s *ListWaitingChatsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWaitingChatsRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *ListWaitingChatsRequest) SetInstanceId(v string) *ListWaitingChatsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWaitingChatsRequest) SetSkillGroupIdList(v string) *ListWaitingChatsRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListWaitingChatsRequest) Validate() error {
	return dara.Validate(s)
}
