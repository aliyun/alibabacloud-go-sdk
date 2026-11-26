// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessApplicationL7SwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetL7Switch(v *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) *UpdatePrivateAccessApplicationL7SwitchResponseBody
	GetL7Switch() *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch
	SetRequestId(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBody
	GetRequestId() *string
}

type UpdatePrivateAccessApplicationL7SwitchResponseBody struct {
	// The Layer 7 access switch configuration of the internal access application after this update.
	L7Switch *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch `json:"L7Switch,omitempty" xml:"L7Switch,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrivateAccessApplicationL7SwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationL7SwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBody) GetL7Switch() *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	return s.L7Switch
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBody) SetL7Switch(v *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) *UpdatePrivateAccessApplicationL7SwitchResponseBody {
	s.L7Switch = v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBody) SetRequestId(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBody) Validate() error {
	if s.L7Switch != nil {
		if err := s.L7Switch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch struct {
	// The ID of the internal access application.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The device tag marking switch. Valid values:
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
	// The effective port ranges for Layer 7 access, which are the intersection of the ports specified in this request and the port ranges already configured for the internal access application. An empty set is returned when Status is set to **Disabled**.
	PortRanges []*UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The source IP marking switch. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	SrcIpMarkStatus *string `json:"SrcIpMarkStatus,omitempty" xml:"SrcIpMarkStatus,omitempty"`
	// The master switch for the Layer 7 access switch of the internal access application. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The request timeout period in seconds.
	//
	// example:
	//
	// 60
	TimeoutSec *int32 `json:"TimeoutSec,omitempty" xml:"TimeoutSec,omitempty"`
	// The URL access whitelist rules for the application. This parameter takes effect only when Status is set to Enabled. An empty value indicates that URLs and request methods are not restricted. A request passes the URL check only when the request path and the method meet any group of rules at the same time. Otherwise, the request is blocked. Passing the URL check does not allow requests that are already blocked by an access policy.
	UrlAccessPaths []*UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths `json:"UrlAccessPaths,omitempty" xml:"UrlAccessPaths,omitempty" type:"Repeated"`
	// The user marking switch. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	UserMarkStatus *string `json:"UserMarkStatus,omitempty" xml:"UserMarkStatus,omitempty"`
	// The host bypass prevention switch. Valid values:
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

func (s UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetDevTagMarkStatus() *string {
	return s.DevTagMarkStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetDownloadAuditStatus() *string {
	return s.DownloadAuditStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetPortRanges() []*UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges {
	return s.PortRanges
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetSrcIpMarkStatus() *string {
	return s.SrcIpMarkStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetStatus() *string {
	return s.Status
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetTimeoutSec() *int32 {
	return s.TimeoutSec
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetUrlAccessPaths() []*UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths {
	return s.UrlAccessPaths
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetUserMarkStatus() *string {
	return s.UserMarkStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) GetZeroTrustStatus() *string {
	return s.ZeroTrustStatus
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetApplicationId(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetDevTagMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.DevTagMarkStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetDownloadAuditStatus(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.DownloadAuditStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetPortRanges(v []*UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.PortRanges = v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetSrcIpMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.SrcIpMarkStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetStatus(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetTimeoutSec(v int32) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.TimeoutSec = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetUrlAccessPaths(v []*UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.UrlAccessPaths = v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetUserMarkStatus(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.UserMarkStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) SetZeroTrustStatus(v string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch {
	s.ZeroTrustStatus = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7Switch) Validate() error {
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UrlAccessPaths != nil {
		for _, item := range s.UrlAccessPaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges struct {
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

func (s UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges) SetBegin(v int32) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges {
	s.Begin = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges) SetEnd(v int32) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges {
	s.End = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchPortRanges) Validate() error {
	return dara.Validate(s)
}

type UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths struct {
	// The list of HTTP request methods allowed for the paths in this group. An empty value indicates that request methods are not restricted.
	AllowMethods []*string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty" type:"Repeated"`
	// The list of URL paths allowed for access in this group. You must specify at least 1 path and up to 20 paths. The paths match the path part in the request, excluding the protocol, domain name, and query parameters.
	Urls []*string `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Repeated"`
}

func (s UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths) GetAllowMethods() []*string {
	return s.AllowMethods
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths) GetUrls() []*string {
	return s.Urls
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths) SetAllowMethods(v []*string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths {
	s.AllowMethods = v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths) SetUrls(v []*string) *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths {
	s.Urls = v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponseBodyL7SwitchUrlAccessPaths) Validate() error {
	return dara.Validate(s)
}
