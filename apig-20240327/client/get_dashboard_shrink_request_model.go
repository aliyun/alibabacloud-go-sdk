// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDashboardShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetDashboardShrinkRequest
	GetAcceptLanguage() *string
	SetApiId(v string) *GetDashboardShrinkRequest
	GetApiId() *string
	SetFilterShrink(v string) *GetDashboardShrinkRequest
	GetFilterShrink() *string
	SetName(v string) *GetDashboardShrinkRequest
	GetName() *string
	SetPluginClassId(v string) *GetDashboardShrinkRequest
	GetPluginClassId() *string
	SetPluginId(v string) *GetDashboardShrinkRequest
	GetPluginId() *string
	SetRouteId(v string) *GetDashboardShrinkRequest
	GetRouteId() *string
	SetSource(v string) *GetDashboardShrinkRequest
	GetSource() *string
	SetUpstreamCluster(v string) *GetDashboardShrinkRequest
	GetUpstreamCluster() *string
}

type GetDashboardShrinkRequest struct {
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
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
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

func (s GetDashboardShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDashboardShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDashboardShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetDashboardShrinkRequest) GetApiId() *string {
	return s.ApiId
}

func (s *GetDashboardShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetDashboardShrinkRequest) GetName() *string {
	return s.Name
}

func (s *GetDashboardShrinkRequest) GetPluginClassId() *string {
	return s.PluginClassId
}

func (s *GetDashboardShrinkRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *GetDashboardShrinkRequest) GetRouteId() *string {
	return s.RouteId
}

func (s *GetDashboardShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *GetDashboardShrinkRequest) GetUpstreamCluster() *string {
	return s.UpstreamCluster
}

func (s *GetDashboardShrinkRequest) SetAcceptLanguage(v string) *GetDashboardShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetApiId(v string) *GetDashboardShrinkRequest {
	s.ApiId = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetFilterShrink(v string) *GetDashboardShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetName(v string) *GetDashboardShrinkRequest {
	s.Name = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetPluginClassId(v string) *GetDashboardShrinkRequest {
	s.PluginClassId = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetPluginId(v string) *GetDashboardShrinkRequest {
	s.PluginId = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetRouteId(v string) *GetDashboardShrinkRequest {
	s.RouteId = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetSource(v string) *GetDashboardShrinkRequest {
	s.Source = &v
	return s
}

func (s *GetDashboardShrinkRequest) SetUpstreamCluster(v string) *GetDashboardShrinkRequest {
	s.UpstreamCluster = &v
	return s
}

func (s *GetDashboardShrinkRequest) Validate() error {
	return dara.Validate(s)
}
