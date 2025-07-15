// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntranetDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CreateIntranetDomainResponseBody
	GetDomainName() *string
	SetRequestId(v string) *CreateIntranetDomainResponseBody
	GetRequestId() *string
}

type CreateIntranetDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// api.demo.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// auditing
	//
	// example:
	//
	// 20D942A5-EDC6-5DA3-93F9-257888399E22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIntranetDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIntranetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntranetDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateIntranetDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIntranetDomainResponseBody) SetDomainName(v string) *CreateIntranetDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *CreateIntranetDomainResponseBody) SetRequestId(v string) *CreateIntranetDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntranetDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
