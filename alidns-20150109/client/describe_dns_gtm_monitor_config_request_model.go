// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmMonitorConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDnsGtmMonitorConfigRequest
	GetLang() *string
	SetMonitorConfigId(v string) *DescribeDnsGtmMonitorConfigRequest
	GetMonitorConfigId() *string
}

type DescribeDnsGtmMonitorConfigRequest struct {
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
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
}

func (s DescribeDnsGtmMonitorConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmMonitorConfigRequest) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeDnsGtmMonitorConfigRequest) SetLang(v string) *DescribeDnsGtmMonitorConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigRequest) SetMonitorConfigId(v string) *DescribeDnsGtmMonitorConfigRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigRequest) Validate() error {
	return dara.Validate(s)
}
