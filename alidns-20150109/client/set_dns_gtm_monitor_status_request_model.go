// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDnsGtmMonitorStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SetDnsGtmMonitorStatusRequest
	GetLang() *string
	SetMonitorConfigId(v string) *SetDnsGtmMonitorStatusRequest
	GetMonitorConfigId() *string
	SetStatus(v string) *SetDnsGtmMonitorStatusRequest
	GetStatus() *string
}

type SetDnsGtmMonitorStatusRequest struct {
	// The language of the values for specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the health check task.
	//
	// This parameter is required.
	//
	// example:
	//
	// MonitorConfigId1
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- OPEN: enables the health check feature.
	//
	// 	- CLOSE: disables the health check feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDnsGtmMonitorStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDnsGtmMonitorStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDnsGtmMonitorStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *SetDnsGtmMonitorStatusRequest) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *SetDnsGtmMonitorStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetDnsGtmMonitorStatusRequest) SetLang(v string) *SetDnsGtmMonitorStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetDnsGtmMonitorStatusRequest) SetMonitorConfigId(v string) *SetDnsGtmMonitorStatusRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *SetDnsGtmMonitorStatusRequest) SetStatus(v string) *SetDnsGtmMonitorStatusRequest {
	s.Status = &v
	return s
}

func (s *SetDnsGtmMonitorStatusRequest) Validate() error {
	return dara.Validate(s)
}
