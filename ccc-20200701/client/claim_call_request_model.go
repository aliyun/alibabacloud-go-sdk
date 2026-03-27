// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClaimCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCandidateUserListJson(v string) *ClaimCallRequest
	GetCandidateUserListJson() *string
	SetInstanceId(v string) *ClaimCallRequest
	GetInstanceId() *string
	SetJobId(v string) *ClaimCallRequest
	GetJobId() *string
	SetSkillGroupId(v string) *ClaimCallRequest
	GetSkillGroupId() *string
	SetTags(v string) *ClaimCallRequest
	GetTags() *string
	SetUserId(v string) *ClaimCallRequest
	GetUserId() *string
}

type ClaimCallRequest struct {
	// example:
	//
	// [
	//
	// {
	//
	// "f0": "zeren001@report-test-2",
	//
	// "f1": "desktop-voip-box@report-test-2"
	//
	// }
	//
	// ]
	CandidateUserListJson *string `json:"CandidateUserListJson,omitempty" xml:"CandidateUserListJson,omitempty"`
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
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// test_sg_****@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// tags
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// invoker@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ClaimCallRequest) String() string {
	return dara.Prettify(s)
}

func (s ClaimCallRequest) GoString() string {
	return s.String()
}

func (s *ClaimCallRequest) GetCandidateUserListJson() *string {
	return s.CandidateUserListJson
}

func (s *ClaimCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ClaimCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *ClaimCallRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ClaimCallRequest) GetTags() *string {
	return s.Tags
}

func (s *ClaimCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *ClaimCallRequest) SetCandidateUserListJson(v string) *ClaimCallRequest {
	s.CandidateUserListJson = &v
	return s
}

func (s *ClaimCallRequest) SetInstanceId(v string) *ClaimCallRequest {
	s.InstanceId = &v
	return s
}

func (s *ClaimCallRequest) SetJobId(v string) *ClaimCallRequest {
	s.JobId = &v
	return s
}

func (s *ClaimCallRequest) SetSkillGroupId(v string) *ClaimCallRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ClaimCallRequest) SetTags(v string) *ClaimCallRequest {
	s.Tags = &v
	return s
}

func (s *ClaimCallRequest) SetUserId(v string) *ClaimCallRequest {
	s.UserId = &v
	return s
}

func (s *ClaimCallRequest) Validate() error {
	return dara.Validate(s)
}
