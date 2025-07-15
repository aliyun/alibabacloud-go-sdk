// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainDetail(v *DescribeLiveDomainDetailResponseBodyDomainDetail) *DescribeLiveDomainDetailResponseBody
	GetDomainDetail() *DescribeLiveDomainDetailResponseBodyDomainDetail
	SetRequestId(v string) *DescribeLiveDomainDetailResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainDetailResponseBody struct {
	// The configuration details of the domain name.
	DomainDetail *DescribeLiveDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 09ABE829-6CD3-4FE0-AFEE-556113E29727
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainDetailResponseBody) GetDomainDetail() *DescribeLiveDomainDetailResponseBodyDomainDetail {
	return s.DomainDetail
}

func (s *DescribeLiveDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainDetailResponseBody) SetDomainDetail(v *DescribeLiveDomainDetailResponseBodyDomainDetail) *DescribeLiveDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeLiveDomainDetailResponseBody) SetRequestId(v string) *DescribeLiveDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainDetailResponseBodyDomainDetail struct {
	// The name of the certificate.
	//
	// example:
	//
	// liveCert****
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The CNAME that is assigned to the domain name. You must add a CNAME record in the system of your Domain Name System (DNS) service provider to map the domain name to the CNAME.
	//
	// >  A time-to-live (TTL) value is specified in the CNAME record of a domain name to indicate how long the CNAME record can be cached on the DNS resolver. If you modify the CNAME record of the domain name, the new settings take effect after the cache expires, which takes about 10 minutes. For more information, see [CNAME resolution](https://help.aliyun.com/document_detail/362010.html).
	//
	// example:
	//
	// learn.developer.aliyundoc.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The description of the domain name.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The streaming domain or ingest domain.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **online**: The domain name is enabled.
	//
	// 	- **offline**: The domain name is disabled.
	//
	// 	- **configuring**: The domain is being configured.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the domain name was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-27T06:51:25Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the domain name was last modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-07T06:51Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- **liveVideo**: streaming domain
	//
	// 	- **liveEdge**: ingest domain
	//
	// example:
	//
	// liveVideo
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty"`
	// The ID of the region where the domain name resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2ogvt4nwmi7i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the SSL certificate is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// Public Key
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// The acceleration region. Valid values:
	//
	// 	- **domestic**: regions in the Chinese mainland.
	//
	// 	- **overseas**: regions outside the Chinese mainland.
	//
	// 	- **global**: regions in and outside the Chinese mainland.
	//
	// example:
	//
	// domestic
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribeLiveDomainDetailResponseBodyDomainDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetCertName() *string {
	return s.CertName
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetCname() *string {
	return s.Cname
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetLiveDomainType() *string {
	return s.LiveDomainType
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) GetScope() *string {
	return s.Scope
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetCertName(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.CertName = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetLiveDomainType(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.LiveDomainType = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetRegion(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetResourceGroupId(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeLiveDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeLiveDomainDetailResponseBodyDomainDetail) Validate() error {
	return dara.Validate(s)
}
