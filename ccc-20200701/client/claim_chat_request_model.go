// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ClaimChatRequest
	GetInstanceId() *string
	SetJobId(v string) *ClaimChatRequest
	GetJobId() *string
	SetSkillGroupId(v string) *ClaimChatRequest
	GetSkillGroupId() *string
}

type ClaimChatRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chat-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ClaimChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ClaimChatRequest) GoString() string {
	return s.String()
}

func (s *ClaimChatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ClaimChatRequest) GetJobId() *string {
	return s.JobId
}

func (s *ClaimChatRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ClaimChatRequest) SetInstanceId(v string) *ClaimChatRequest {
	s.InstanceId = &v
	return s
}

func (s *ClaimChatRequest) SetJobId(v string) *ClaimChatRequest {
	s.JobId = &v
	return s
}

func (s *ClaimChatRequest) SetSkillGroupId(v string) *ClaimChatRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ClaimChatRequest) Validate() error {
	return dara.Validate(s)
}
