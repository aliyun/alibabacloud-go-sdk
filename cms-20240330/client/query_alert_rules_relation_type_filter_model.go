// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesRelationTypeFilter interface {
	dara.Model
	String() string
	GoString() string
	SetIn(v []*string) *QueryAlertRulesRelationTypeFilter
	GetIn() []*string
	SetNotIn(v []*string) *QueryAlertRulesRelationTypeFilter
	GetNotIn() []*string
}

type QueryAlertRulesRelationTypeFilter struct {
	In    []*string `json:"in,omitempty" xml:"in,omitempty" type:"Repeated"`
	NotIn []*string `json:"notIn,omitempty" xml:"notIn,omitempty" type:"Repeated"`
}

func (s QueryAlertRulesRelationTypeFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesRelationTypeFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesRelationTypeFilter) GetIn() []*string {
	return s.In
}

func (s *QueryAlertRulesRelationTypeFilter) GetNotIn() []*string {
	return s.NotIn
}

func (s *QueryAlertRulesRelationTypeFilter) SetIn(v []*string) *QueryAlertRulesRelationTypeFilter {
	s.In = v
	return s
}

func (s *QueryAlertRulesRelationTypeFilter) SetNotIn(v []*string) *QueryAlertRulesRelationTypeFilter {
	s.NotIn = v
	return s
}

func (s *QueryAlertRulesRelationTypeFilter) Validate() error {
	return dara.Validate(s)
}
