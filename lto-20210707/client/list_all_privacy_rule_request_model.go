// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllPrivacyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAllPrivacyRuleRequest
	GetRegionId() *string
}

type ListAllPrivacyRuleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllPrivacyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllPrivacyRuleRequest) GoString() string {
	return s.String()
}

func (s *ListAllPrivacyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllPrivacyRuleRequest) SetRegionId(v string) *ListAllPrivacyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllPrivacyRuleRequest) Validate() error {
	return dara.Validate(s)
}
