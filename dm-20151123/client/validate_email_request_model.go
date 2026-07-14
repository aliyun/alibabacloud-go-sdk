// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckGraylist(v bool) *ValidateEmailRequest
	GetCheckGraylist() *bool
	SetEmail(v string) *ValidateEmailRequest
	GetEmail() *string
	SetProbeType(v string) *ValidateEmailRequest
	GetProbeType() *string
	SetTimeout(v int64) *ValidateEmailRequest
	GetTimeout() *int64
}

type ValidateEmailRequest struct {
	// Specifies whether to check the graylist. Default value: false. The result is asynchronously notified through EventBridge.
	//
	// example:
	//
	// true
	CheckGraylist *bool `json:"CheckGraylist,omitempty" xml:"CheckGraylist,omitempty"`
	// The email address to validate.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx@yyy.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The probe type. Valid values:
	//
	// - FULL: Enables all detection capabilities, including SMTP probing. Because SMTP probing involves remote connections, the overall latency is high. This value is suitable for scenarios that are not sensitive to response time. Each detection consumes 1 address validation quota.
	//
	// - BASIC_ONLY: Enables all detection capabilities except SMTP probing, with low latency. This value is suitable for scenarios that are sensitive to response time, such as real-time validation during registration to check whether an email address is a disposable mailbox or an MX-forwarded abnormal address, to prevent batch registration by the cyber underground economy chain. Each detection consumes 1/3 of an address validation quota.
	//
	// example:
	//
	// FULL
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	// The timeout period. Default value: 60 seconds.
	//
	// example:
	//
	// 20
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ValidateEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateEmailRequest) GoString() string {
	return s.String()
}

func (s *ValidateEmailRequest) GetCheckGraylist() *bool {
	return s.CheckGraylist
}

func (s *ValidateEmailRequest) GetEmail() *string {
	return s.Email
}

func (s *ValidateEmailRequest) GetProbeType() *string {
	return s.ProbeType
}

func (s *ValidateEmailRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *ValidateEmailRequest) SetCheckGraylist(v bool) *ValidateEmailRequest {
	s.CheckGraylist = &v
	return s
}

func (s *ValidateEmailRequest) SetEmail(v string) *ValidateEmailRequest {
	s.Email = &v
	return s
}

func (s *ValidateEmailRequest) SetProbeType(v string) *ValidateEmailRequest {
	s.ProbeType = &v
	return s
}

func (s *ValidateEmailRequest) SetTimeout(v int64) *ValidateEmailRequest {
	s.Timeout = &v
	return s
}

func (s *ValidateEmailRequest) Validate() error {
	return dara.Validate(s)
}
