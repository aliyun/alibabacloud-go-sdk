// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCmsExporterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *InstallCmsExporterRequest
	GetClusterId() *string
	SetCmsArgs(v string) *InstallCmsExporterRequest
	GetCmsArgs() *string
	SetDirectArgs(v string) *InstallCmsExporterRequest
	GetDirectArgs() *string
	SetEnableTag(v bool) *InstallCmsExporterRequest
	GetEnableTag() *bool
	SetRegionId(v string) *InstallCmsExporterRequest
	GetRegionId() *string
}

type InstallCmsExporterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cloud services that you want to monitor. The CmsArgs parameter is the startup parameter of the cms-exporter collector. Separate multiple cloud services with number signs (`#`).
	//
	// example:
	//
	// hologres#cen
	CmsArgs *string `json:"CmsArgs,omitempty" xml:"CmsArgs,omitempty"`
	// The recently monitored cloud services. Separate multiple cloud services with number signs (`#`).
	//
	// example:
	//
	// hologres#cen
	DirectArgs *string `json:"DirectArgs,omitempty" xml:"DirectArgs,omitempty"`
	// Specifies whether to collect the aliyun tags attached to each cloud service. Default value: false.
	//
	// example:
	//
	// false
	EnableTag *bool `json:"EnableTag,omitempty" xml:"EnableTag,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCmsExporterRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCmsExporterRequest) GoString() string {
	return s.String()
}

func (s *InstallCmsExporterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallCmsExporterRequest) GetCmsArgs() *string {
	return s.CmsArgs
}

func (s *InstallCmsExporterRequest) GetDirectArgs() *string {
	return s.DirectArgs
}

func (s *InstallCmsExporterRequest) GetEnableTag() *bool {
	return s.EnableTag
}

func (s *InstallCmsExporterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallCmsExporterRequest) SetClusterId(v string) *InstallCmsExporterRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallCmsExporterRequest) SetCmsArgs(v string) *InstallCmsExporterRequest {
	s.CmsArgs = &v
	return s
}

func (s *InstallCmsExporterRequest) SetDirectArgs(v string) *InstallCmsExporterRequest {
	s.DirectArgs = &v
	return s
}

func (s *InstallCmsExporterRequest) SetEnableTag(v bool) *InstallCmsExporterRequest {
	s.EnableTag = &v
	return s
}

func (s *InstallCmsExporterRequest) SetRegionId(v string) *InstallCmsExporterRequest {
	s.RegionId = &v
	return s
}

func (s *InstallCmsExporterRequest) Validate() error {
	return dara.Validate(s)
}
