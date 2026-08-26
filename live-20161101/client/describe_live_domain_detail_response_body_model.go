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
	// The details of the domain configuration.
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
	if s.DomainDetail != nil {
		if err := s.DomainDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainDetailResponseBodyDomainDetail struct {
	// The name of the certificate.
	//
	// example:
	//
	// liveCert****
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The canonical name (CNAME). A CNAME is generated for the live streaming domain. You must add a CNAME record at your DNS provider to map the live streaming domain to this CNAME.
	//
	// > Local DNS records are cached. After you add the CNAME record, it may take up to 10 minutes to take effect. For more information, see [FAQ about CNAME records](https://help.aliyun.com/document_detail/362010.html).
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
	// The ingest domain or streaming domain.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of the domain. Valid values:
	//
	// - **online**: enabled.
	//
	// - **offline**: disabled.
	//
	// - **configuring**: being configured.
	//
	// example:
	//
	// online
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// The time when the domain was added. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-27T06:51:25Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the domain was last modified. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-08-07T06:51Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The type of the domain name. Valid values:
	//
	// - **liveVideo**: streaming domain.
	//
	// - **liveEdge**: ingest domain.
	//
	// example:
	//
	// liveVideo
	LiveDomainType *string `json:"LiveDomainType,omitempty" xml:"LiveDomainType,omitempty"`
	// The region where the domain name is added.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-**k3bpq2yjw22**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether SSL is enabled. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
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
	// - **domestic**: the Chinese mainland.
	//
	// - **overseas**: regions outside the Chinese mainland.
	//
	// - **global**: global.
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
