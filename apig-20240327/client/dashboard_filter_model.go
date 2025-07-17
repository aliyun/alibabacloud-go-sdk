// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashboardFilter interface {
	dara.Model
	String() string
	GoString() string
	SetRouteName(v string) *DashboardFilter
	GetRouteName() *string
}

type DashboardFilter struct {
	// example:
	//
	// test
	RouteName *string `json:"routeName,omitempty" xml:"routeName,omitempty"`
}

func (s DashboardFilter) String() string {
	return dara.Prettify(s)
}

func (s DashboardFilter) GoString() string {
	return s.String()
}

func (s *DashboardFilter) GetRouteName() *string {
	return s.RouteName
}

func (s *DashboardFilter) SetRouteName(v string) *DashboardFilter {
	s.RouteName = &v
	return s
}

func (s *DashboardFilter) Validate() error {
	return dara.Validate(s)
}
