// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMainDomainNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLevel(v int32) *GetMainDomainNameResponseBody
	GetDomainLevel() *int32
	SetMainDomainName(v string) *GetMainDomainNameResponseBody
	GetMainDomainName() *string
	SetRR(v string) *GetMainDomainNameResponseBody
	GetRR() *string
	SetRequestId(v string) *GetMainDomainNameResponseBody
	GetRequestId() *string
}

type GetMainDomainNameResponseBody struct {
	// The level of the domain name.
	//
	// example:
	//
	// 2
	DomainLevel *int32 `json:"DomainLevel,omitempty" xml:"DomainLevel,omitempty"`
	// The root domain name.
	//
	// example:
	//
	// example.com
	MainDomainName *string `json:"MainDomainName,omitempty" xml:"MainDomainName,omitempty"`
	// The host record.
	//
	// example:
	//
	// sub
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMainDomainNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMainDomainNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetMainDomainNameResponseBody) GetDomainLevel() *int32 {
	return s.DomainLevel
}

func (s *GetMainDomainNameResponseBody) GetMainDomainName() *string {
	return s.MainDomainName
}

func (s *GetMainDomainNameResponseBody) GetRR() *string {
	return s.RR
}

func (s *GetMainDomainNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMainDomainNameResponseBody) SetDomainLevel(v int32) *GetMainDomainNameResponseBody {
	s.DomainLevel = &v
	return s
}

func (s *GetMainDomainNameResponseBody) SetMainDomainName(v string) *GetMainDomainNameResponseBody {
	s.MainDomainName = &v
	return s
}

func (s *GetMainDomainNameResponseBody) SetRR(v string) *GetMainDomainNameResponseBody {
	s.RR = &v
	return s
}

func (s *GetMainDomainNameResponseBody) SetRequestId(v string) *GetMainDomainNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMainDomainNameResponseBody) Validate() error {
	return dara.Validate(s)
}
