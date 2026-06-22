// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListHoneypotProbeRequest
	GetCurrentPage() *int32
	SetDisplayName(v string) *ListHoneypotProbeRequest
	GetDisplayName() *string
	SetLang(v string) *ListHoneypotProbeRequest
	GetLang() *string
	SetPageSize(v int32) *ListHoneypotProbeRequest
	GetPageSize() *int32
	SetProbeStatus(v string) *ListHoneypotProbeRequest
	GetProbeStatus() *string
	SetProbeType(v string) *ListHoneypotProbeRequest
	GetProbeType() *string
}

type ListHoneypotProbeRequest struct {
	// The page number of the page to return. Minimum value: **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the probe.
	//
	// example:
	//
	// probe-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page in a paged query. Default value: 20. If you leave this parameter empty, 20 entries are returned per page during paging.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The probe status. Valid values:
	//
	// - **installed**: Installation succeeded.
	//
	// - **install_failed**: Installation failed.
	//
	// - **online**: Normal.
	//
	// - **offline**: Offline.
	//
	// - **unnormal**: Service exception.
	//
	// - **unprobe**: Unauthorized.
	//
	// - **uninstalling**: Uninstalling.
	//
	// - **uninstalled**: Uninstallation succeeded.
	//
	// - **uninstall_failed**: Uninstallation failed.
	//
	// - **not_exist**: Not installed.
	//
	// example:
	//
	// online
	ProbeStatus *string `json:"ProbeStatus,omitempty" xml:"ProbeStatus,omitempty"`
	// The probe type. Valid values:
	//
	// - **host_probe**: host probe
	//
	// - **vpc_black_hole_probe**: VPC blackhole probe.
	//
	// example:
	//
	// host_probe
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
}

func (s ListHoneypotProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotProbeRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListHoneypotProbeRequest) GetLang() *string {
	return s.Lang
}

func (s *ListHoneypotProbeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotProbeRequest) GetProbeStatus() *string {
	return s.ProbeStatus
}

func (s *ListHoneypotProbeRequest) GetProbeType() *string {
	return s.ProbeType
}

func (s *ListHoneypotProbeRequest) SetCurrentPage(v int32) *ListHoneypotProbeRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotProbeRequest) SetDisplayName(v string) *ListHoneypotProbeRequest {
	s.DisplayName = &v
	return s
}

func (s *ListHoneypotProbeRequest) SetLang(v string) *ListHoneypotProbeRequest {
	s.Lang = &v
	return s
}

func (s *ListHoneypotProbeRequest) SetPageSize(v int32) *ListHoneypotProbeRequest {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotProbeRequest) SetProbeStatus(v string) *ListHoneypotProbeRequest {
	s.ProbeStatus = &v
	return s
}

func (s *ListHoneypotProbeRequest) SetProbeType(v string) *ListHoneypotProbeRequest {
	s.ProbeType = &v
	return s
}

func (s *ListHoneypotProbeRequest) Validate() error {
	return dara.Validate(s)
}
