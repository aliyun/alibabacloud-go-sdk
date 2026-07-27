// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesEntityTypeFilter interface {
	dara.Model
	String() string
	GoString() string
	SetIn(v []*string) *QueryAlertRulesEntityTypeFilter
	GetIn() []*string
	SetNotIn(v []*string) *QueryAlertRulesEntityTypeFilter
	GetNotIn() []*string
}

type QueryAlertRulesEntityTypeFilter struct {
	In    []*string `json:"in,omitempty" xml:"in,omitempty" type:"Repeated"`
	NotIn []*string `json:"notIn,omitempty" xml:"notIn,omitempty" type:"Repeated"`
}

func (s QueryAlertRulesEntityTypeFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesEntityTypeFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesEntityTypeFilter) GetIn() []*string {
	return s.In
}

func (s *QueryAlertRulesEntityTypeFilter) GetNotIn() []*string {
	return s.NotIn
}

func (s *QueryAlertRulesEntityTypeFilter) SetIn(v []*string) *QueryAlertRulesEntityTypeFilter {
	s.In = v
	return s
}

func (s *QueryAlertRulesEntityTypeFilter) SetNotIn(v []*string) *QueryAlertRulesEntityTypeFilter {
	s.NotIn = v
	return s
}

func (s *QueryAlertRulesEntityTypeFilter) Validate() error {
	return dara.Validate(s)
}
