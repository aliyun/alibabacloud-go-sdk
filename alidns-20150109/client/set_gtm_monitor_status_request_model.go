// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGtmMonitorStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SetGtmMonitorStatusRequest
	GetLang() *string
	SetMonitorConfigId(v string) *SetGtmMonitorStatusRequest
	GetMonitorConfigId() *string
	SetStatus(v string) *SetGtmMonitorStatusRequest
	GetStatus() *string
}

type SetGtmMonitorStatusRequest struct {
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The health check ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc1234
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// Specifies whether health check is enabled for the address pool. Valid values:
	//
	// 	- **OPEN**: Enabled
	//
	// 	- **CLOSE**: Disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetGtmMonitorStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetGtmMonitorStatusRequest) GoString() string {
	return s.String()
}

func (s *SetGtmMonitorStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *SetGtmMonitorStatusRequest) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *SetGtmMonitorStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetGtmMonitorStatusRequest) SetLang(v string) *SetGtmMonitorStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetGtmMonitorStatusRequest) SetMonitorConfigId(v string) *SetGtmMonitorStatusRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *SetGtmMonitorStatusRequest) SetStatus(v string) *SetGtmMonitorStatusRequest {
	s.Status = &v
	return s
}

func (s *SetGtmMonitorStatusRequest) Validate() error {
	return dara.Validate(s)
}
