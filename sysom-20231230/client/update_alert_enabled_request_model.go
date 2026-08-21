// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *UpdateAlertEnabledRequest
	GetXDebugId() *string
	SetEnabled(v bool) *UpdateAlertEnabledRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateAlertEnabledRequest
	GetId() *int64
	SetXSysomInvokeSource(v string) *UpdateAlertEnabledRequest
	GetXSysomInvokeSource() *string
}

type UpdateAlertEnabledRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// Specifies whether the alert policy is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The ID of the alert policy.
	//
	// example:
	//
	// 1
	Id                 *int64  `json:"id,omitempty" xml:"id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s UpdateAlertEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertEnabledRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertEnabledRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *UpdateAlertEnabledRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAlertEnabledRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAlertEnabledRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *UpdateAlertEnabledRequest) SetXDebugId(v string) *UpdateAlertEnabledRequest {
	s.XDebugId = &v
	return s
}

func (s *UpdateAlertEnabledRequest) SetEnabled(v bool) *UpdateAlertEnabledRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateAlertEnabledRequest) SetId(v int64) *UpdateAlertEnabledRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlertEnabledRequest) SetXSysomInvokeSource(v string) *UpdateAlertEnabledRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *UpdateAlertEnabledRequest) Validate() error {
	return dara.Validate(s)
}
