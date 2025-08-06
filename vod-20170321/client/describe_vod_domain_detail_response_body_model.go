// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainDetail(v *DescribeVodDomainDetailResponseBodyDomainDetail) *DescribeVodDomainDetailResponseBody
	GetDomainDetail() *DescribeVodDomainDetailResponseBodyDomainDetail
	SetRequestId(v string) *DescribeVodDomainDetailResponseBody
	GetRequestId() *string
}

type DescribeVodDomainDetailResponseBody struct {
	// The basic information about the domain name for CDN.
	DomainDetail *DescribeVodDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 09ABE829-6CD3-4FE0-556113E2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBody) GetDomainDetail() *DescribeVodDomainDetailResponseBodyDomainDetail {
	return s.DomainDetail
}

func (s *DescribeVodDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainDetailResponseBody) SetDomainDetail(v *DescribeVodDomainDetailResponseBodyDomainDetail) *DescribeVodDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeVodDomainDetailResponseBody) SetRequestId(v string) *DescribeVodDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainDetailResponseBodyDomainDetail struct {
	// The name of the certificate. The value of this parameter is returned if HTTPS is enabled.
	//
	// example:
	//
	// testCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The CNAME that is assigned to the domain name for CDN. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name for CDN to the CNAME.
	//
	// example:
	//
	// example.com.w.alikunlun.net
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The description of the domain name for CDN.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name for CDN.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the domain name for CDN. Value values:
	//
	// 	- **online**: indicates that the domain name is enabled.
	//
	// 	- **offline**: indicates that the domain name is disabled.
	//
	// 	- **configuring**: indicates that the domain name is being configured.
	//
	// 	- **configure_failed**: indicates that the domain name failed to be configured.
	//
	// 	- **checking**: indicates that the domain name is under review.
	//
	// 	- **check_failed**: indicates that the domain name failed the review.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the domain name for CDN was added. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-27T06:51:26Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The last time when the domain name for CDN was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-27T06:55:26Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the Secure Sockets Layer (SSL) certificate is enabled. Valid values:
	//
	// 	- **on**: indicates that the SSL certificate is enabled.
	//
	// 	- **off**: indicates that the SSL certificate is disabled.
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The public key of the certificate. The value of this parameter is returned if HTTPS is enabled.
	//
	// example:
	//
	// yourSSLPub
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// This parameter is applicable to users of level 3 or higher in mainland China and users outside mainland China. Valid values:
	//
	// 	- **domestic**: mainland China. This is the default value.
	//
	// 	- **overseas**: outside mainland China.
	//
	// 	- **global**: regions in and outside mainland China.
	//
	// example:
	//
	// domestic
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The information about the origin server.
	Sources *DescribeVodDomainDetailResponseBodyDomainDetailSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	// The weight of the origin server.
	//
	// example:
	//
	// 1
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetCname() *string {
	return s.Cname
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetScope() *string {
	return s.Scope
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetSources() *DescribeVodDomainDetailResponseBodyDomainDetailSources {
	return s.Sources
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) GetWeight() *string {
	return s.Weight
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetCertName(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.CertName = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSources(v *DescribeVodDomainDetailResponseBodyDomainDetailSources) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Sources = v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetWeight(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Weight = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainDetailResponseBodyDomainDetailSources struct {
	Source []*DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSources) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSources) GetSource() []*DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	return s.Source
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSources) SetSource(v []*DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) *DescribeVodDomainDetailResponseBodyDomainDetailSources {
	s.Source = v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSources) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// ****.oss-cn-hangzhou.aliyuncs.com
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The status of the origin server. Valid values:
	//
	// 	- **online**: indicates that the origin server is enabled.
	//
	// 	- **offline**: indicates that the origin server is disabled.
	//
	// example:
	//
	// online
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The port number. Valid values: 443 and 80.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority of the origin server.
	//
	// example:
	//
	// 50
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the origin server. Valid values:
	//
	// 	- **ipaddr**: a server that you can access by using an IP address.
	//
	// 	- **domain**: a server that you can access by using a domain name.
	//
	// 	- **oss**: the URL of an Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin server if multiple origin servers have been specified.
	//
	// example:
	//
	// 10
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetContent(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetEnabled(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Enabled = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetPort(v int32) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetPriority(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetType(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetWeight(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) Validate() error {
	return dara.Validate(s)
}
