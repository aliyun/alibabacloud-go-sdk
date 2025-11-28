// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivacyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivacyRuleId(v string) *DeletePrivacyRuleRequest
	GetPrivacyRuleId() *string
	SetRegionId(v string) *DeletePrivacyRuleRequest
	GetRegionId() *string
}

type DeletePrivacyRuleRequest struct {
	// This parameter is required.
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePrivacyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivacyRuleRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivacyRuleRequest) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *DeletePrivacyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePrivacyRuleRequest) SetPrivacyRuleId(v string) *DeletePrivacyRuleRequest {
	s.PrivacyRuleId = &v
	return s
}

func (s *DeletePrivacyRuleRequest) SetRegionId(v string) *DeletePrivacyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePrivacyRuleRequest) Validate() error {
	return dara.Validate(s)
}
