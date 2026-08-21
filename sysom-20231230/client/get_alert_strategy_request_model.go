// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetAlertStrategyRequest
	GetXDebugId() *string
	SetId(v int64) *GetAlertStrategyRequest
	GetId() *int64
	SetXSysomInvokeSource(v string) *GetAlertStrategyRequest
	GetXSysomInvokeSource() *string
}

type GetAlertStrategyRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 1
	Id                 *int64  `json:"id,omitempty" xml:"id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetAlertStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetAlertStrategyRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetAlertStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *GetAlertStrategyRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetAlertStrategyRequest) SetXDebugId(v string) *GetAlertStrategyRequest {
	s.XDebugId = &v
	return s
}

func (s *GetAlertStrategyRequest) SetId(v int64) *GetAlertStrategyRequest {
	s.Id = &v
	return s
}

func (s *GetAlertStrategyRequest) SetXSysomInvokeSource(v string) *GetAlertStrategyRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetAlertStrategyRequest) Validate() error {
	return dara.Validate(s)
}
