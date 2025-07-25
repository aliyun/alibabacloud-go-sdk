// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTxtRecordForVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *GetTxtRecordForVerifyRequest
	GetDomainName() *string
	SetLang(v string) *GetTxtRecordForVerifyRequest
	GetLang() *string
	SetType(v string) *GetTxtRecordForVerifyRequest
	GetType() *string
}

type GetTxtRecordForVerifyRequest struct {
	// The ID of the Domain Name System (DNS) record. You can call the [DescribeDomainRecords](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomainrecords?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The feature verified by using the TXT record. Valid values:
	//
	// 	- ADD_SUB_DOMAIN: add a subdomain for verification.
	//
	// 	- RETRIEVAL: add other subdomains for verification
	//
	// This parameter is required.
	//
	// example:
	//
	// ADD_SUB_DOMAIN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTxtRecordForVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTxtRecordForVerifyRequest) GoString() string {
	return s.String()
}

func (s *GetTxtRecordForVerifyRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *GetTxtRecordForVerifyRequest) GetLang() *string {
	return s.Lang
}

func (s *GetTxtRecordForVerifyRequest) GetType() *string {
	return s.Type
}

func (s *GetTxtRecordForVerifyRequest) SetDomainName(v string) *GetTxtRecordForVerifyRequest {
	s.DomainName = &v
	return s
}

func (s *GetTxtRecordForVerifyRequest) SetLang(v string) *GetTxtRecordForVerifyRequest {
	s.Lang = &v
	return s
}

func (s *GetTxtRecordForVerifyRequest) SetType(v string) *GetTxtRecordForVerifyRequest {
	s.Type = &v
	return s
}

func (s *GetTxtRecordForVerifyRequest) Validate() error {
	return dara.Validate(s)
}
