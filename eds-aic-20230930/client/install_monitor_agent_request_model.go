// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMonitorAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *InstallMonitorAgentRequest
	GetAndroidInstanceIds() []*string
	SetSaleMode(v string) *InstallMonitorAgentRequest
	GetSaleMode() *string
}

type InstallMonitorAgentRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// NODE
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
}

func (s InstallMonitorAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallMonitorAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallMonitorAgentRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *InstallMonitorAgentRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *InstallMonitorAgentRequest) SetAndroidInstanceIds(v []*string) *InstallMonitorAgentRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *InstallMonitorAgentRequest) SetSaleMode(v string) *InstallMonitorAgentRequest {
	s.SaleMode = &v
	return s
}

func (s *InstallMonitorAgentRequest) Validate() error {
	return dara.Validate(s)
}
