// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDiagnosticInterruptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *SendDiagnosticInterruptRequest
	GetDryRun() *bool
	SetInstanceId(v string) *SendDiagnosticInterruptRequest
	GetInstanceId() *string
}

type SendDiagnosticInterruptRequest struct {
	// Specifies whether to only precheck the request. Valid values:
	//
	// - true: Sends a check request without sending the NMI command.
	//
	// - false: Sends a normal NMI request to trigger a crash dump. Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID to which you want to send a diagnostic break.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SendDiagnosticInterruptRequest) String() string {
	return dara.Prettify(s)
}

func (s SendDiagnosticInterruptRequest) GoString() string {
	return s.String()
}

func (s *SendDiagnosticInterruptRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *SendDiagnosticInterruptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendDiagnosticInterruptRequest) SetDryRun(v bool) *SendDiagnosticInterruptRequest {
	s.DryRun = &v
	return s
}

func (s *SendDiagnosticInterruptRequest) SetInstanceId(v string) *SendDiagnosticInterruptRequest {
	s.InstanceId = &v
	return s
}

func (s *SendDiagnosticInterruptRequest) Validate() error {
	return dara.Validate(s)
}
