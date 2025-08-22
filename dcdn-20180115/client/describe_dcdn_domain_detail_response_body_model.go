// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainDetail(v *DescribeDcdnDomainDetailResponseBodyDomainDetail) *DescribeDcdnDomainDetailResponseBody
	GetDomainDetail() *DescribeDcdnDomainDetailResponseBodyDomainDetail
	SetRequestId(v string) *DescribeDcdnDomainDetailResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainDetailResponseBody struct {
	// The information about the accelerated domain name.
	DomainDetail *DescribeDcdnDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 09ABE829-6CD3-4FE0-AFEE-556113E29727
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponseBody) GetDomainDetail() *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	return s.DomainDetail
}

func (s *DescribeDcdnDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainDetailResponseBody) SetDomainDetail(v *DescribeDcdnDomainDetailResponseBodyDomainDetail) *DescribeDcdnDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBody) SetRequestId(v string) *DescribeDcdnDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainDetailResponseBodyDomainDetail struct {
	// The CNAME that is assigned to the accelerated domain name. You must add the CNAME record to the system of your Domain Name System (DNS) provider to map the accelerated domain name to the CNAME.
	//
	// example:
	//
	// example.aliyundoc.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The information about the Internet content provider (ICP) filing of the domain name.
	//
	// example:
	//
	// Beijing ICP Filing No. 1703xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the accelerated domain name. Valid values:
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// 	- **configuring**
	//
	// 	- **configure_failed**
	//
	// 	- **checking**
	//
	// 	- **check_failed**
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// Computing service type. Valid values:
	//
	// 	- **routine**
	//
	// 	- **image**
	//
	// 	- **cloudFunction**
	//
	// example:
	//
	// routine
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The time when the domain name was added.
	//
	// example:
	//
	// 2017-11-27T06:51:26Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the domain name was last modified.
	//
	// example:
	//
	// 2017-11-27T06:51:25Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyuji4b6r4**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the Security Socket Layer (SSL) certificate is enabled. Valid values:
	//
	// 	- **on**: **enabled**
	//
	// 	- **off**: **disabled**
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The public key of the certificate if HTTPS is enabled.
	//
	// example:
	//
	// xxx
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// Acceleration scenario. Valid values:
	//
	// 	- **apiscene**: API acceleration.
	//
	// 	- **webservicescene**: website acceleration.
	//
	// 	- **staticscene**: video, image, and text acceleration.
	//
	// 	- **an empty string**: no acceleration scenario is used.
	//
	// example:
	//
	// apiscene
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The acceleration region. Default value: domestic. Valid values:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: global (excluding the Chinese mainland)
	//
	// 	- **global**: global
	//
	// example:
	//
	// overseas
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The information about the origin server.
	Sources *DescribeDcdnDomainDetailResponseBodyDomainDetailSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetCname() *string {
	return s.Cname
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetFunctionType() *string {
	return s.FunctionType
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetScene() *string {
	return s.Scene
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetScope() *string {
	return s.Scope
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) GetSources() *DescribeDcdnDomainDetailResponseBodyDomainDetailSources {
	return s.Sources
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetFunctionType(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.FunctionType = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetResourceGroupId(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetScene(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Scene = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) SetSources(v *DescribeDcdnDomainDetailResponseBodyDomainDetailSources) *DescribeDcdnDomainDetailResponseBodyDomainDetail {
	s.Sources = v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainDetailResponseBodyDomainDetailSources struct {
	Source []*DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetailSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetailSources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSources) GetSource() []*DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	return s.Source
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSources) SetSource(v []*DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) *DescribeDcdnDomainDetailResponseBodyDomainDetailSources {
	s.Source = v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSources) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource struct {
	// The address of the origin server.
	//
	// example:
	//
	// example.org
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The status.
	//
	// example:
	//
	// online
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The port over which requests are redirected to the origin server. Ports 443 and 80 are supported.
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
	// 	- **ipaddr**: an IP address
	//
	// 	- **domain**: an origin domain name
	//
	// 	- **oss**: the domain name of an Object Storage Service (OSS) bucket
	//
	// example:
	//
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the origin server if multiple origin servers are specified.
	//
	// example:
	//
	// 20
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) GetEnabled() *string {
	return s.Enabled
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) GetPriority() *string {
	return s.Priority
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) GetWeight() *string {
	return s.Weight
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetContent(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetEnabled(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Enabled = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetPort(v int32) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetPriority(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetType(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Type = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) SetWeight(v string) *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Weight = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponseBodyDomainDetailSourcesSource) Validate() error {
	return dara.Validate(s)
}
