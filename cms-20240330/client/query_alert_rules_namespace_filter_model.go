// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesNamespaceFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *QueryAlertRulesNamespaceFilter
	GetEq() *string
}

type QueryAlertRulesNamespaceFilter struct {
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s QueryAlertRulesNamespaceFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesNamespaceFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesNamespaceFilter) GetEq() *string {
	return s.Eq
}

func (s *QueryAlertRulesNamespaceFilter) SetEq(v string) *QueryAlertRulesNamespaceFilter {
	s.Eq = &v
	return s
}

func (s *QueryAlertRulesNamespaceFilter) Validate() error {
	return dara.Validate(s)
}
