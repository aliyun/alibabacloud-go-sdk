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
	// 路由表：自定义域名访问时的 PATH 到 Function 的映射列表。
	//
	// example:
	//
	// []
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
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
