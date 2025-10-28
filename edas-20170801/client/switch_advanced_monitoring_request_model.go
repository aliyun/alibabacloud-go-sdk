// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchAdvancedMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SwitchAdvancedMonitoringRequest
	GetAppId() *string
	SetEnableAdvancedMonitoring(v bool) *SwitchAdvancedMonitoringRequest
	GetEnableAdvancedMonitoring() *bool
}

type SwitchAdvancedMonitoringRequest struct {
	// The ID of the application for which you want to query or configure the advanced application monitoring feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9e224bc6-a646-4484-****-e617b7e7****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to enable the advanced application monitoring feature. Valid values:
	//
	// 	- true: enables the advanced application monitoring feature.
	//
	// 	- false: disables the advanced application monitoring feature.
	//
	// If you call this operation to query the status of the advanced application monitoring feature, you do not need to specify this parameter.
	//
	// example:
	//
	// true
	EnableAdvancedMonitoring *bool `json:"EnableAdvancedMonitoring,omitempty" xml:"EnableAdvancedMonitoring,omitempty"`
}

func (s SwitchAdvancedMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchAdvancedMonitoringRequest) GoString() string {
	return s.String()
}

func (s *SwitchAdvancedMonitoringRequest) GetAppId() *string {
	return s.AppId
}

func (s *SwitchAdvancedMonitoringRequest) GetEnableAdvancedMonitoring() *bool {
	return s.EnableAdvancedMonitoring
}

func (s *SwitchAdvancedMonitoringRequest) SetAppId(v string) *SwitchAdvancedMonitoringRequest {
	s.AppId = &v
	return s
}

func (s *SwitchAdvancedMonitoringRequest) SetEnableAdvancedMonitoring(v bool) *SwitchAdvancedMonitoringRequest {
	s.EnableAdvancedMonitoring = &v
	return s
}

func (s *SwitchAdvancedMonitoringRequest) Validate() error {
	return dara.Validate(s)
}
