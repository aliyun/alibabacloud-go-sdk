// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleLabelFilter interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v map[string]*string) *AlertRuleLabelFilter
	GetLabels() map[string]*string
	SetOpt(v string) *AlertRuleLabelFilter
	GetOpt() *string
}

type AlertRuleLabelFilter struct {
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	Opt    *string            `json:"opt,omitempty" xml:"opt,omitempty"`
}

func (s AlertRuleLabelFilter) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleLabelFilter) GoString() string {
	return s.String()
}

func (s *AlertRuleLabelFilter) GetLabels() map[string]*string {
	return s.Labels
}

func (s *AlertRuleLabelFilter) GetOpt() *string {
	return s.Opt
}

func (s *AlertRuleLabelFilter) SetLabels(v map[string]*string) *AlertRuleLabelFilter {
	s.Labels = v
	return s
}

func (s *AlertRuleLabelFilter) SetOpt(v string) *AlertRuleLabelFilter {
	s.Opt = &v
	return s
}

func (s *AlertRuleLabelFilter) Validate() error {
	return dara.Validate(s)
}
