// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRouteConfig interface {
	dara.Model
	String() string
	GoString() string
	SetRoutes(v []*PathConfig) *RouteConfig
	GetRoutes() []*PathConfig
}

type RouteConfig struct {
	Routes []*PathConfig `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s RouteConfig) String() string {
	return dara.Prettify(s)
}

func (s RouteConfig) GoString() string {
	return s.String()
}

func (s *RouteConfig) GetRoutes() []*PathConfig {
	return s.Routes
}

func (s *RouteConfig) SetRoutes(v []*PathConfig) *RouteConfig {
	s.Routes = v
	return s
}

func (s *RouteConfig) Validate() error {
	return dara.Validate(s)
}
