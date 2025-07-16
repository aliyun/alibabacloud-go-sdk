// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDomains(v []*ListGatewayDomainsResponseBodyCustomDomains) *ListGatewayDomainsResponseBody
	GetCustomDomains() []*ListGatewayDomainsResponseBodyCustomDomains
	SetMessage(v string) *ListGatewayDomainsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayDomainsResponseBody
	GetRequestId() *string
}

type ListGatewayDomainsResponseBody struct {
	// The custom domain names.
	CustomDomains []*ListGatewayDomainsResponseBodyCustomDomains `json:"CustomDomains,omitempty" xml:"CustomDomains,omitempty" type:"Repeated"`
	// The message that is returned.
	//
	// example:
	//
	// Successfully get custom domains
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGatewayDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainsResponseBody) GetCustomDomains() []*ListGatewayDomainsResponseBodyCustomDomains {
	return s.CustomDomains
}

func (s *ListGatewayDomainsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayDomainsResponseBody) SetCustomDomains(v []*ListGatewayDomainsResponseBodyCustomDomains) *ListGatewayDomainsResponseBody {
	s.CustomDomains = v
	return s
}

func (s *ListGatewayDomainsResponseBody) SetMessage(v string) *ListGatewayDomainsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayDomainsResponseBody) SetRequestId(v string) *ListGatewayDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGatewayDomainsResponseBodyCustomDomains struct {
	CertificateEndDate *string `json:"CertificateEndDate,omitempty" xml:"CertificateEndDate,omitempty"`
	// The ID of the SSL certificate bound to the domain name. Obtain the certificate ID after you upload or purchase a certificate in the [Certificate Management Service](https://yundunnext.console.aliyun.com/?spm=5176.2020520163.console-base_help.2.4b3baJixaJixOc\\&p=cas) console.
	//
	// example:
	//
	// 1473**25
	CertificateId        *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName      *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	CertificateStartDate *string `json:"CertificateStartDate,omitempty" xml:"CertificateStartDate,omitempty"`
	CertificateStatus    *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The custom domain name.
	//
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The domain name type.
	//
	// Valid value:
	//
	// 	- intranet: internal network.
	//
	// 	- internet: public network.
	//
	// example:
	//
	// intranet
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGatewayDomainsResponseBodyCustomDomains) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainsResponseBodyCustomDomains) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetCertificateEndDate() *string {
	return s.CertificateEndDate
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetCertificateName() *string {
	return s.CertificateName
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetCertificateStartDate() *string {
	return s.CertificateStartDate
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetDomain() *string {
	return s.Domain
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetType() *string {
	return s.Type
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetCertificateEndDate(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.CertificateEndDate = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetCertificateId(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.CertificateId = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetCertificateName(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.CertificateName = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetCertificateStartDate(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.CertificateStartDate = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetCertificateStatus(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.CertificateStatus = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetCreateTime(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.CreateTime = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetDomain(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.Domain = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetType(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.Type = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) SetUpdateTime(v string) *ListGatewayDomainsResponseBodyCustomDomains {
	s.UpdateTime = &v
	return s
}

func (s *ListGatewayDomainsResponseBodyCustomDomains) Validate() error {
	return dara.Validate(s)
}
