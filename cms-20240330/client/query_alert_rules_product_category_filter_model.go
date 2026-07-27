// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesProductCategoryFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *QueryAlertRulesProductCategoryFilter
	GetEq() *string
}

type QueryAlertRulesProductCategoryFilter struct {
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s QueryAlertRulesProductCategoryFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesProductCategoryFilter) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesProductCategoryFilter) GetEq() *string {
	return s.Eq
}

func (s *QueryAlertRulesProductCategoryFilter) SetEq(v string) *QueryAlertRulesProductCategoryFilter {
	s.Eq = &v
	return s
}

func (s *QueryAlertRulesProductCategoryFilter) Validate() error {
	return dara.Validate(s)
}
