// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDashboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetDashboardRequest
	GetAcceptLanguage() *string
	SetApiId(v string) *GetDashboardRequest
	GetApiId() *string
	SetFilter(v *GetDashboardRequestFilter) *GetDashboardRequest
	GetFilter() *GetDashboardRequestFilter
	SetName(v string) *GetDashboardRequest
	GetName() *string
	SetPluginClassId(v string) *GetDashboardRequest
	GetPluginClassId() *string
	SetPluginId(v string) *GetDashboardRequest
	GetPluginId() *string
	SetRouteId(v string) *GetDashboardRequest
	GetRouteId() *string
	SetSource(v string) *GetDashboardRequest
	GetSource() *string
	SetUpstreamCluster(v string) *GetDashboardRequest
	GetUpstreamCluster() *string
}

type GetDashboardRequest struct {
	// The language. Valid values: zh (Chinese) and en (English).
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	// API ID
	//
	// example:
	//
	// api-c9uuekzmia8q2****
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// The filter configurations.
	Filter *GetDashboardRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// The dashboard name.
	//
	// 	- LOG: access logs
	//
	// 	- PLUGIN: plug-in logs
	//
	// example:
	//
	// PLUGIN
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The plug-in ID.
	//
	// example:
	//
	// pls-dn82a9djd8z****
	PluginClassId *string `json:"pluginClassId,omitempty" xml:"pluginClassId,omitempty"`
	PluginId      *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	RouteId       *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
	// The dashboard source. Valid values:
	//
	// 	- SLS: Simple Log Service
	//
	// example:
	//
	// SLS
	Source          *string `json:"source,omitempty" xml:"source,omitempty"`
	UpstreamCluster *string `json:"upstreamCluster,omitempty" xml:"upstreamCluster,omitempty"`
}

func (s GetDashboardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDashboardRequest) GoString() string {
	return s.String()
}

func (s *GetDashboardRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetDashboardRequest) GetApiId() *string {
	return s.ApiId
}

func (s *GetDashboardRequest) GetFilter() *GetDashboardRequestFilter {
	return s.Filter
}

func (s *GetDashboardRequest) GetName() *string {
	return s.Name
}

func (s *GetDashboardRequest) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *GetDashboardRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *GetDashboardRequest) GetRouteId() *string {
	return s.RouteId
}

func (s *GetDashboardRequest) GetSource() *string {
	return s.Source
}

func (s *GetDashboardRequest) GetUpstreamCluster() *string {
	return s.UpstreamCluster
}

func (s *GetDashboardRequest) SetAcceptLanguage(v string) *GetDashboardRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetDashboardRequest) SetApiId(v string) *GetDashboardRequest {
	s.ApiId = &v
	return s
}

func (s *GetDashboardRequest) SetFilter(v *GetDashboardRequestFilter) *GetDashboardRequest {
	s.Filter = v
	return s
}

func (s *GetDashboardRequest) SetName(v string) *GetDashboardRequest {
	s.Name = &v
	return s
}

func (s *GetDashboardRequest) SetPluginClassId(v string) *GetDashboardRequest {
	s.PluginClassId = &v
	return s
}

func (s *GetDashboardRequest) SetPluginId(v string) *GetDashboardRequest {
	s.PluginId = &v
	return s
}

func (s *GetDashboardRequest) SetRouteId(v string) *GetDashboardRequest {
	s.RouteId = &v
	return s
}

func (s *GetDashboardRequest) SetSource(v string) *GetDashboardRequest {
	s.Source = &v
	return s
}

func (s *GetDashboardRequest) SetUpstreamCluster(v string) *GetDashboardRequest {
	s.UpstreamCluster = &v
	return s
}

func (s *GetDashboardRequest) Validate() error {
	return dara.Validate(s)
}

type GetDashboardRequestFilter struct {
	// The route name.
	//
	// example:
	//
	// test-route
	RouteName *string `json:"routeName,omitempty" xml:"routeName,omitempty"`
}

func (s GetDashboardRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetDashboardRequestFilter) GoString() string {
	return s.String()
}

func (s *GetDashboardRequestFilter) GetRouteName() *string {
	return s.RouteName
}

func (s *GetDashboardRequestFilter) SetRouteName(v string) *GetDashboardRequestFilter {
	s.RouteName = &v
	return s
}

func (s *GetDashboardRequestFilter) Validate() error {
	return dara.Validate(s)
}
