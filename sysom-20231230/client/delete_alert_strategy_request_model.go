// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *DeleteAlertStrategyRequest
	GetXDebugId() *string
	SetId(v string) *DeleteAlertStrategyRequest
	GetId() *string
	SetXSysomInvokeSource(v string) *DeleteAlertStrategyRequest
	GetXSysomInvokeSource() *string
}

type DeleteAlertStrategyRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The ID of the alert policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id                 *string `json:"id,omitempty" xml:"id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s DeleteAlertStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertStrategyRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *DeleteAlertStrategyRequest) GetId() *string {
	return s.Id
}

func (s *DeleteAlertStrategyRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *DeleteAlertStrategyRequest) SetXDebugId(v string) *DeleteAlertStrategyRequest {
	s.XDebugId = &v
	return s
}

func (s *DeleteAlertStrategyRequest) SetId(v string) *DeleteAlertStrategyRequest {
	s.Id = &v
	return s
}

func (s *DeleteAlertStrategyRequest) SetXSysomInvokeSource(v string) *DeleteAlertStrategyRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *DeleteAlertStrategyRequest) Validate() error {
	return dara.Validate(s)
}
