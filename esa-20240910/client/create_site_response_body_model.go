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
	// The nameservers assigned by ESA. The values are separated by commas (,). This parameter is returned if you set AccessType to NS. In this case, you must change the nameservers of your domain to the assigned ones. Then, you can verify the domain ownership and activate your website.
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
	// The website ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The verification code for the website. If you set AccessType to CNAME, you need to add a TXT record whose hostname is **_esaauth.[websiteDomainName]*	- and record value is the value of VerifyCode to the DNS records of your domain. ****Then, you can verify the domain ownership and activate your website.
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
