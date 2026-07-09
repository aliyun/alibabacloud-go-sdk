// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasourceTypeFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *DatasourceTypeFilter
	GetEq() *string
}

type DatasourceTypeFilter struct {
	// Specifies the data source type for an exact match.
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s DatasourceTypeFilter) String() string {
	return dara.Prettify(s)
}

func (s DatasourceTypeFilter) GoString() string {
	return s.String()
}

func (s *DatasourceTypeFilter) GetEq() *string {
	return s.Eq
}

func (s *DatasourceTypeFilter) SetEq(v string) *DatasourceTypeFilter {
	s.Eq = &v
	return s
}

func (s *DatasourceTypeFilter) Validate() error {
	return dara.Validate(s)
}
