// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesResourcesFilter interface {
	dara.Model
	String() string
	GoString() string
	SetContains(v []*string) *QueryAlertRulesResourcesFilter
	GetContains() []*string
	SetNotContains(v []*string) *QueryAlertRulesResourcesFilter
	GetNotContains() []*string
}

type QueryAlertRulesResourcesFilter struct {
	// Matches any item in the list (OR semantics).
	//
	// example:
	//
	// ["i-bp1a2b3c4d5e6f7g8h9i","i-bp9h8g7f6e5d4c3b2a1"]
	Contains []*string `json:"contains,omitempty" xml:"contains,omitempty" type:"Repeated"`
	// The resource instance ID blacklist filter that excludes alert rules whose associated resources contain any instance ID in the array.
	//
	// example:
	//
	// ["i-bp0z9y8x7w6v5u4t3s2"]
	NotContains []*string `json:"notContains,omitempty" xml:"notContains,omitempty" type:"Repeated"`
}

func (s QueryAlertRulesResourcesFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesResourcesFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesResourcesFilter) GetContains() []*string {
	return s.Contains
}

func (s *QueryAlertRulesResourcesFilter) GetNotContains() []*string {
	return s.NotContains
}

func (s *QueryAlertRulesResourcesFilter) SetContains(v []*string) *QueryAlertRulesResourcesFilter {
	s.Contains = v
	return s
}

func (s *QueryAlertRulesResourcesFilter) SetNotContains(v []*string) *QueryAlertRulesResourcesFilter {
	s.NotContains = v
	return s
}

func (s *QueryAlertRulesResourcesFilter) Validate() error {
	return dara.Validate(s)
}
