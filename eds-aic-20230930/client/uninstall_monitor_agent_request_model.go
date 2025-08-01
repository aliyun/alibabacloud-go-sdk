// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallMonitorAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *UninstallMonitorAgentRequest
	GetAndroidInstanceIds() []*string
	SetSaleMode(v string) *UninstallMonitorAgentRequest
	GetSaleMode() *string
}

type UninstallMonitorAgentRequest struct {
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// Node
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
}

func (s UninstallMonitorAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallMonitorAgentRequest) GoString() string {
	return s.String()
}

func (s *UninstallMonitorAgentRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *UninstallMonitorAgentRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *UninstallMonitorAgentRequest) SetAndroidInstanceIds(v []*string) *UninstallMonitorAgentRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *UninstallMonitorAgentRequest) SetSaleMode(v string) *UninstallMonitorAgentRequest {
	s.SaleMode = &v
	return s
}

func (s *UninstallMonitorAgentRequest) Validate() error {
	return dara.Validate(s)
}
