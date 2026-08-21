// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessApplicationL7SwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetL7Switches(v []*ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) *ListPrivateAccessApplicationL7SwitchesResponseBody
	GetL7Switches() []*ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches
	SetRequestId(v string) *ListPrivateAccessApplicationL7SwitchesResponseBody
	GetRequestId() *string
}

type ListPrivateAccessApplicationL7SwitchesResponseBody struct {
	// The Layer 7 application configurations.
	L7Switches []*ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches `json:"L7Switches,omitempty" xml:"L7Switches,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// ED459CAD-8D3F-51B8-AEA5-CAABC0325022
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivateAccessApplicationL7SwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationL7SwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBody) GetL7Switches() []*ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	return s.L7Switches
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBody) SetL7Switches(v []*ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) *ListPrivateAccessApplicationL7SwitchesResponseBody {
	s.L7Switches = v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBody) SetRequestId(v string) *ListPrivateAccessApplicationL7SwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBody) Validate() error {
	if s.L7Switches != nil {
		for _, item := range s.L7Switches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches struct {
	// The ID of the internal-facing application.
	//
	// example:
	//
	// pa-application-bbbc550d7c6e4db6
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The device tag mark switch. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	DevTagMarkStatus *string `json:"DevTagMarkStatus,omitempty" xml:"DevTagMarkStatus,omitempty"`
	// The download audit switch for sensitive applications. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	DownloadAuditStatus *string `json:"DownloadAuditStatus,omitempty" xml:"DownloadAuditStatus,omitempty"`
	// The port ranges.
	PortRanges []*ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The source IP mark switch. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	SrcIpMarkStatus *string `json:"SrcIpMarkStatus,omitempty" xml:"SrcIpMarkStatus,omitempty"`
	// The status of the internal-facing access policy. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The request timeout period.
	//
	// example:
	//
	// 60
	TimeoutSec *int32 `json:"TimeoutSec,omitempty" xml:"TimeoutSec,omitempty"`
	// The user mark switch. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	UserMarkStatus *string `json:"UserMarkStatus,omitempty" xml:"UserMarkStatus,omitempty"`
	// The host bypass prevention switch. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	ZeroTrustStatus *string `json:"ZeroTrustStatus,omitempty" xml:"ZeroTrustStatus,omitempty"`
}

func (s ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetDevTagMarkStatus() *string {
	return s.DevTagMarkStatus
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetDownloadAuditStatus() *string {
	return s.DownloadAuditStatus
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetPortRanges() []*ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges {
	return s.PortRanges
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetSrcIpMarkStatus() *string {
	return s.SrcIpMarkStatus
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetStatus() *string {
	return s.Status
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetTimeoutSec() *int32 {
	return s.TimeoutSec
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetUserMarkStatus() *string {
	return s.UserMarkStatus
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) GetZeroTrustStatus() *string {
	return s.ZeroTrustStatus
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetApplicationId(v string) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.ApplicationId = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetDevTagMarkStatus(v string) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.DevTagMarkStatus = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetDownloadAuditStatus(v string) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.DownloadAuditStatus = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetPortRanges(v []*ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.PortRanges = v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetSrcIpMarkStatus(v string) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.SrcIpMarkStatus = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetStatus(v string) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetTimeoutSec(v int32) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.TimeoutSec = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetUserMarkStatus(v string) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.UserMarkStatus = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) SetZeroTrustStatus(v string) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches {
	s.ZeroTrustStatus = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7Switches) Validate() error {
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

type ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges struct {
	// The start port.
	//
	// example:
	//
	// 2379
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// The end port. The value must be greater than or equal to the start port.
	//
	// example:
	//
	// 24
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges) SetBegin(v int32) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges {
	s.Begin = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges) SetEnd(v int32) *ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges {
	s.End = &v
	return s
}

func (s *ListPrivateAccessApplicationL7SwitchesResponseBodyL7SwitchesPortRanges) Validate() error {
	return dara.Validate(s)
}
