// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsDeleteRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SdnsDeleteRuleCategoryRequest
	GetId() *string
}

type SdnsDeleteRuleCategoryRequest struct {
	// The policy ID of the current account. This parameter is required. You can obtain the ID from the Id field in the response of CreateSdnsRuleCategory or the list returned by SdnsListRuleCategories. Use a decimal string to prevent precision loss of large integers.
	//
	// example:
	//
	// 123456789012345678
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SdnsDeleteRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s SdnsDeleteRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *SdnsDeleteRuleCategoryRequest) GetId() *string {
	return s.Id
}

func (s *SdnsDeleteRuleCategoryRequest) SetId(v string) *SdnsDeleteRuleCategoryRequest {
	s.Id = &v
	return s
}

func (s *SdnsDeleteRuleCategoryRequest) Validate() error {
	return dara.Validate(s)
}
