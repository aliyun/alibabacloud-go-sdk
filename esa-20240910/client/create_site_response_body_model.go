// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNameServerList(v string) *CreateSiteResponseBody
	GetNameServerList() *string
	SetRequestId(v string) *CreateSiteResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *CreateSiteResponseBody
	GetSiteId() *int64
	SetVerifyCode(v string) *CreateSiteResponseBody
	GetVerifyCode() *string
}

type CreateSiteResponseBody struct {
	// The list of name servers assigned to the site, separated by commas (,). This field has a value when the access type of the site is NS. You need to change the DNS servers of the site to these name servers. Then you can verify the site ownership and activate the site.
	//
	// example:
	//
	// ns1.example.com,ns2.example.com
	NameServerList *string `json:"NameServerList,omitempty" xml:"NameServerList,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The verification code of the site. When the access type of the site is CNAME, you need to add a TXT record with the record name **_esaauth.[site name]*	- and the record value set to the **verification code*	- on the DNS server of the site. Then you can verify the site ownership and activate the site.
	//
	// example:
	//
	// verify_aah9dioasmov****
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s CreateSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteResponseBody) GetNameServerList() *string {
	return s.NameServerList
}

func (s *CreateSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSiteResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateSiteResponseBody) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *CreateSiteResponseBody) SetNameServerList(v string) *CreateSiteResponseBody {
	s.NameServerList = &v
	return s
}

func (s *CreateSiteResponseBody) SetRequestId(v string) *CreateSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSiteResponseBody) SetSiteId(v int64) *CreateSiteResponseBody {
	s.SiteId = &v
	return s
}

func (s *CreateSiteResponseBody) SetVerifyCode(v string) *CreateSiteResponseBody {
	s.VerifyCode = &v
	return s
}

func (s *CreateSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
