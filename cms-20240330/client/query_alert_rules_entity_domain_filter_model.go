// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesEntityDomainFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *QueryAlertRulesEntityDomainFilter
	GetEq() *string
}

type QueryAlertRulesEntityDomainFilter struct {
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s QueryAlertRulesEntityDomainFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesEntityDomainFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesEntityDomainFilter) GetEq() *string {
	return s.Eq
}

func (s *QueryAlertRulesEntityDomainFilter) SetEq(v string) *QueryAlertRulesEntityDomainFilter {
	s.Eq = &v
	return s
}

func (s *QueryAlertRulesEntityDomainFilter) Validate() error {
	return dara.Validate(s)
}
