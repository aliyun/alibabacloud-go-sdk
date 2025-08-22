// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainDetail(v *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) *DescribeDcdnIpaDomainDetailResponseBody
	GetDomainDetail() *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail
	SetRequestId(v string) *DescribeDcdnIpaDomainDetailResponseBody
	GetRequestId() *string
}

type DescribeDcdnIpaDomainDetailResponseBody struct {
	// The details about the accelerated domain name.
	DomainDetail *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 09ABE829-6CD3-4FE0-AFEE-556113E29727
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpaDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponseBody) GetDomainDetail() *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	return s.DomainDetail
}

func (s *DescribeDcdnIpaDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnIpaDomainDetailResponseBody) SetDomainDetail(v *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) *DescribeDcdnIpaDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBody) SetRequestId(v string) *DescribeDcdnIpaDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaDomainDetailResponseBodyDomainDetail struct {
	// Indicates the name of the certificate if the HTTPS protocol is enabled.
	//
	// example:
	//
	// yourCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The CNAME assigned to the domain name.
	//
	// example:
	//
	// example.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The description.
	//
	// example:
	//
	// audit failed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the accelerated domain name. Valid values:
	//
	// 	- **online**: enabled
	//
	// 	- **offline**: disabled
	//
	// 	- **configuring**: configuring
	//
	// 	- **configure_failed**: configuration failed
	//
	// 	- **checking**: reviewing
	//
	// 	- **check_failed:*	- review failed
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2017-11-27T06:51:26Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the domain name was last modified.
	//
	// example:
	//
	// 2017-11-27T06:51:26Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyuji4b6xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the Security Socket Layer (SSL) certificate is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**.
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The public key of the certificate if HTTPS is enabled.
	//
	// example:
	//
	// SSLPub
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// The acceleration region. Valid values:
	//
	// 	- domestic: Chinese mainland
	//
	// 	- overseas: outside the Chinese mainland
	//
	// 	- global: global
	//
	// example:
	//
	// overseas
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The information about the origin server.
	Sources *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetCname() *string {
	return s.Cname
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetScope() *string {
	return s.Scope
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) GetSources() *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources {
	return s.Sources
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetCertName(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetResourceGroupId(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) SetSources(v *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail {
	s.Sources = v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources struct {
	Source []*DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) GetSource() []*DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	return s.Source
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) SetSource(v []*DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources {
	s.Source = v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSources) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// xxx.oss-cn-hangzhou.aliyuncs.com
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The status.
	//
	// example:
	//
	// online
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The custom port. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The priority.
	//
	// example:
	//
	// 50
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The type of the origin server. Valid values:
	//
	// 	- **ipaddr**: an origin IP address
	//
	// 	- **domain**: a domain name.
	//
	// 	- **oss**: Object Storage Service (OSS) buckets are not supported.
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

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetContent(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetEnabled(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Enabled = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetPort(v int32) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetPriority(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetType(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) SetWeight(v string) *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponseBodyDomainDetailSourcesSource) Validate() error {
	return dara.Validate(s)
}
