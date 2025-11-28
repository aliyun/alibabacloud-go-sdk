// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivacyRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNum(v int32) *ListPrivacyRuleRequest
	GetNum() *int32
	SetRegionId(v string) *ListPrivacyRuleRequest
	GetRegionId() *string
	SetSize(v int32) *ListPrivacyRuleRequest
	GetSize() *int32
}

type ListPrivacyRuleRequest struct {
	// This parameter is required.
	Num      *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListPrivacyRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleRequest) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListPrivacyRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrivacyRuleRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListPrivacyRuleRequest) SetNum(v int32) *ListPrivacyRuleRequest {
	s.Num = &v
	return s
}

func (s *ListPrivacyRuleRequest) SetRegionId(v string) *ListPrivacyRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrivacyRuleRequest) SetSize(v int32) *ListPrivacyRuleRequest {
	s.Size = &v
	return s
}

func (s *ListPrivacyRuleRequest) Validate() error {
	return dara.Validate(s)
}
