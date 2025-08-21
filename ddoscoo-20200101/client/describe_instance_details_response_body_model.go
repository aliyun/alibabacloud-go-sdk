// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceDetails(v []*DescribeInstanceDetailsResponseBodyInstanceDetails) *DescribeInstanceDetailsResponseBody
	GetInstanceDetails() []*DescribeInstanceDetailsResponseBodyInstanceDetails
	SetRequestId(v string) *DescribeInstanceDetailsResponseBody
	GetRequestId() *string
}

type DescribeInstanceDetailsResponseBody struct {
	// The IP address and ISP line information about the Anti-DDoS Proxy instance.
	InstanceDetails []*DescribeInstanceDetailsResponseBodyInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3C814429-21A5-4673-827E-FDD19DC75681
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBody) GetInstanceDetails() []*DescribeInstanceDetailsResponseBodyInstanceDetails {
	return s.InstanceDetails
}

func (s *DescribeInstanceDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceDetailsResponseBody) SetInstanceDetails(v []*DescribeInstanceDetailsResponseBodyInstanceDetails) *DescribeInstanceDetailsResponseBody {
	s.InstanceDetails = v
	return s
}

func (s *DescribeInstanceDetailsResponseBody) SetRequestId(v string) *DescribeInstanceDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceDetailsResponseBodyInstanceDetails struct {
	// The IP address information about the Anti-DDoS Proxy instance.
	EipInfos []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos `json:"EipInfos,omitempty" xml:"EipInfos,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-zvp2eibz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The protection line of the instance.
	//
	// example:
	//
	// coop-line-001
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) GetEipInfos() []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	return s.EipInfos
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) GetLine() *string {
	return s.Line
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetEipInfos(v []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.EipInfos = v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetInstanceId(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetLine(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.Line = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos struct {
	// Indicates whether a custom certificate is configured.
	//
	// example:
	//
	// true
	CertConfigured *bool `json:"CertConfigured,omitempty" xml:"CertConfigured,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.117.XX.XX
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// default
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// The IP address-based forwarding mode of the instance. Valid values:
	//
	// 	- **fnat**: Requests from IPv4 addresses are forwarded to origin servers that use IPv4 addresses and requests from IPv6 addresses are forwarded to origin servers that use IPv6 addresses.
	//
	// 	- **v6tov4**: All requests are forwarded to origin servers that use IPv4 addresses.
	//
	// example:
	//
	// fnat
	IpMode *string `json:"IpMode,omitempty" xml:"IpMode,omitempty"`
	// The IP version of the protocol. Valid values:
	//
	// 	- **Ipv4**: IPv4
	//
	// 	- **Ipv6**: IPv6
	//
	// example:
	//
	// Ipv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether the TLS 1.3 version is supported.
	//
	// example:
	//
	// false
	Ssl13Enabled *bool `json:"Ssl13Enabled,omitempty" xml:"Ssl13Enabled,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **normal**: indicates that the instance is normal.
	//
	// 	- **expired**: indicates that the instance expired.
	//
	// 	- **defense**: indicates that traffic scrubbing is performed on the asset that is protected by the instance.
	//
	// 	- **blackhole**: indicates that blackhole filtering is triggered for the asset that is protected by the instance.
	//
	// 	- **punished**: indicates that the instance is in penalty.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Transport Layer Security (TLS) version that is configured.
	//
	// example:
	//
	// tls1.2
	TlsVersion *string `json:"TlsVersion,omitempty" xml:"TlsVersion,omitempty"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GetCertConfigured() *bool {
	return s.CertConfigured
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GetEip() *string {
	return s.Eip
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GetFunctionVersion() *string {
	return s.FunctionVersion
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GetIpMode() *string {
	return s.IpMode
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GetSsl13Enabled() *bool {
	return s.Ssl13Enabled
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) GetTlsVersion() *string {
	return s.TlsVersion
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetCertConfigured(v bool) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.CertConfigured = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetEip(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.Eip = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetFunctionVersion(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.FunctionVersion = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetIpMode(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.IpMode = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetIpVersion(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetSsl13Enabled(v bool) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.Ssl13Enabled = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetStatus(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.Status = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) SetTlsVersion(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos {
	s.TlsVersion = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfos) Validate() error {
	return dara.Validate(s)
}
