// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDingTalkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDingTalkRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeDingTalkRequest
	GetPageSize() *int32
	SetRuleActionName(v string) *DescribeDingTalkRequest
	GetRuleActionName() *string
}

type DescribeDingTalkRequest struct {
	// The number of the page to return.Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the notification.
	//
	// example:
	//
	// Vulnerability notification
	RuleActionName *string `json:"RuleActionName,omitempty" xml:"RuleActionName,omitempty"`
}

func (s DescribeDingTalkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDingTalkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDingTalkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDingTalkRequest) GetRuleActionName() *string {
	return s.RuleActionName
}

func (s *DescribeDingTalkRequest) SetCurrentPage(v int32) *DescribeDingTalkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDingTalkRequest) SetPageSize(v int32) *DescribeDingTalkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDingTalkRequest) SetRuleActionName(v string) *DescribeDingTalkRequest {
	s.RuleActionName = &v
	return s
}

func (s *DescribeDingTalkRequest) Validate() error {
	return dara.Validate(s)
}
