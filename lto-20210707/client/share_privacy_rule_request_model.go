// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSharePrivacyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberIdList(v string) *SharePrivacyRuleRequest
	GetMemberIdList() *string
	SetPrivacyRuleId(v string) *SharePrivacyRuleRequest
	GetPrivacyRuleId() *string
	SetRegionId(v string) *SharePrivacyRuleRequest
	GetRegionId() *string
}

type SharePrivacyRuleRequest struct {
	MemberIdList *string `json:"MemberIdList,omitempty" xml:"MemberIdList,omitempty"`
	// This parameter is required.
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SharePrivacyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s SharePrivacyRuleRequest) GoString() string {
	return s.String()
}

func (s *SharePrivacyRuleRequest) GetMemberIdList() *string {
	return s.MemberIdList
}

func (s *SharePrivacyRuleRequest) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *SharePrivacyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SharePrivacyRuleRequest) SetMemberIdList(v string) *SharePrivacyRuleRequest {
	s.MemberIdList = &v
	return s
}

func (s *SharePrivacyRuleRequest) SetPrivacyRuleId(v string) *SharePrivacyRuleRequest {
	s.PrivacyRuleId = &v
	return s
}

func (s *SharePrivacyRuleRequest) SetRegionId(v string) *SharePrivacyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *SharePrivacyRuleRequest) Validate() error {
	return dara.Validate(s)
}
