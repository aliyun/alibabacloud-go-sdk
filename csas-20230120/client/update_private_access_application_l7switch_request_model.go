// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessApplicationL7SwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetApplicationId() *string
	SetDevTagMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetDevTagMarkStatus() *string
	SetDownloadAuditStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetDownloadAuditStatus() *string
	SetPortRanges(v []*UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetPortRanges() []*UpdatePrivateAccessApplicationL7SwitchRequestPortRanges
	SetSrcIpMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetSrcIpMarkStatus() *string
	SetStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetStatus() *string
	SetTimeoutSec(v int32) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetTimeoutSec() *int32
	SetUserMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetUserMarkStatus() *string
	SetZeroTrustStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest
	GetZeroTrustStatus() *string
}

type UpdatePrivateAccessApplicationL7SwitchRequest struct {
	// The ID of the internal-facing application. Required.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The device tag mark switch. Required. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	DevTagMarkStatus *string `json:"DevTagMarkStatus,omitempty" xml:"DevTagMarkStatus,omitempty"`
	// The sensitive application download audit switch. Optional. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	DownloadAuditStatus *string `json:"DownloadAuditStatus,omitempty" xml:"DownloadAuditStatus,omitempty"`
	// The collection of port ranges for the internal-facing application. Multiple port ranges cannot be duplicated or overlap. You can specify up to 50 port ranges. This parameter takes effect and is validated only when Status is set to **Enabled**. If this parameter is not specified or an empty collection is passed in, the default ports 80, 443, 8080, and 465 are used. The effective ports are the intersection of the ports specified in this request and the port ranges already configured for the internal-facing application.
	PortRanges []*UpdatePrivateAccessApplicationL7SwitchRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The source IP mark switch. Required. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	SrcIpMarkStatus *string `json:"SrcIpMarkStatus,omitempty" xml:"SrcIpMarkStatus,omitempty"`
	// The master switch for Layer 7 access of the internal-facing application. Required. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// When the value is **Disabled**, PortRanges is neither validated nor saved.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The request timeout period, in seconds. Valid values: 1 to 3600. Default value: 60. If this parameter is not specified or an invalid value is specified, the value 60 is used.
	//
	// example:
	//
	// 60
	TimeoutSec *int32 `json:"TimeoutSec,omitempty" xml:"TimeoutSec,omitempty"`
	// The user mark switch. Required. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	UserMarkStatus *string `json:"UserMarkStatus,omitempty" xml:"UserMarkStatus,omitempty"`
	// The host bypass prevention switch. Required. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	ZeroTrustStatus *string `json:"ZeroTrustStatus,omitempty" xml:"ZeroTrustStatus,omitempty"`
}

func (s UpdatePrivateAccessApplicationL7SwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationL7SwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetDevTagMarkStatus() *string {
	return s.DevTagMarkStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetDownloadAuditStatus() *string {
	return s.DownloadAuditStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetPortRanges() []*UpdatePrivateAccessApplicationL7SwitchRequestPortRanges {
	return s.PortRanges
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetSrcIpMarkStatus() *string {
	return s.SrcIpMarkStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetTimeoutSec() *int32 {
	return s.TimeoutSec
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetUserMarkStatus() *string {
	return s.UserMarkStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) GetZeroTrustStatus() *string {
	return s.ZeroTrustStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetApplicationId(v string) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetDevTagMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.DevTagMarkStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetDownloadAuditStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.DownloadAuditStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetPortRanges(v []*UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.PortRanges = v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetSrcIpMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.SrcIpMarkStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetTimeoutSec(v int32) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.TimeoutSec = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetUserMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.UserMarkStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) SetZeroTrustStatus(v string) *UpdatePrivateAccessApplicationL7SwitchRequest {
	s.ZeroTrustStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequest) Validate() error {
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePrivateAccessApplicationL7SwitchRequestPortRanges struct {
	// The start port. The value must be less than or equal to the end port.
	//
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// The end port. The value must be greater than or equal to the start port.
	//
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) SetBegin(v int32) *UpdatePrivateAccessApplicationL7SwitchRequestPortRanges {
	s.Begin = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) SetEnd(v int32) *UpdatePrivateAccessApplicationL7SwitchRequestPortRanges {
	s.End = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchRequestPortRanges) Validate() error {
	return dara.Validate(s)
}
