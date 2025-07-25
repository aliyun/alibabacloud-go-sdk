// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmMonitorConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGtmMonitorConfigRequest
	GetLang() *string
	SetMonitorConfigId(v string) *DescribeGtmMonitorConfigRequest
	GetMonitorConfigId() *string
}

type DescribeGtmMonitorConfigRequest struct {
	// The language of the values of specific response parameters.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the health check configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
}

func (s DescribeGtmMonitorConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmMonitorConfigRequest) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeGtmMonitorConfigRequest) SetLang(v string) *DescribeGtmMonitorConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmMonitorConfigRequest) SetMonitorConfigId(v string) *DescribeGtmMonitorConfigRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeGtmMonitorConfigRequest) Validate() error {
	return dara.Validate(s)
}
