// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrometheusNamedQueryEntry interface {
	dara.Model
	String() string
	GoString() string
	SetExpr(v string) *PrometheusNamedQueryEntry
	GetExpr() *string
	SetName(v string) *PrometheusNamedQueryEntry
	GetName() *string
}

type PrometheusNamedQueryEntry struct {
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PrometheusNamedQueryEntry) String() string {
	return dara.Prettify(s)
}

func (s PrometheusNamedQueryEntry) GoString() string {
	return s.String()
}

func (s *PrometheusNamedQueryEntry) GetExpr() *string {
	return s.Expr
}

func (s *PrometheusNamedQueryEntry) GetName() *string {
	return s.Name
}

func (s *PrometheusNamedQueryEntry) SetExpr(v string) *PrometheusNamedQueryEntry {
	s.Expr = &v
	return s
}

func (s *PrometheusNamedQueryEntry) SetName(v string) *PrometheusNamedQueryEntry {
	s.Name = &v
	return s
}

func (s *PrometheusNamedQueryEntry) Validate() error {
	return dara.Validate(s)
}
