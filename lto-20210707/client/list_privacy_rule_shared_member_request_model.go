// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivacyRuleSharedMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivacyRuleId(v string) *ListPrivacyRuleSharedMemberRequest
	GetPrivacyRuleId() *string
	SetRegionId(v string) *ListPrivacyRuleSharedMemberRequest
	GetRegionId() *string
}

type ListPrivacyRuleSharedMemberRequest struct {
	// This parameter is required.
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPrivacyRuleSharedMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleSharedMemberRequest) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleSharedMemberRequest) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *ListPrivacyRuleSharedMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrivacyRuleSharedMemberRequest) SetPrivacyRuleId(v string) *ListPrivacyRuleSharedMemberRequest {
	s.PrivacyRuleId = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberRequest) SetRegionId(v string) *ListPrivacyRuleSharedMemberRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberRequest) Validate() error {
	return dara.Validate(s)
}
