// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMainDomainNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainLevel(v int64) *GetMainDomainNameResponseBody
	GetDomainLevel() *int64
	SetDomainName(v string) *GetMainDomainNameResponseBody
	GetDomainName() *string
	SetRR(v string) *GetMainDomainNameResponseBody
	GetRR() *string
	SetRequestId(v string) *GetMainDomainNameResponseBody
	GetRequestId() *string
}

type GetMainDomainNameResponseBody struct {
	// The level of the entered domain name.
	//
	// example:
	//
	// 2
	DomainLevel *int64 `json:"DomainLevel,omitempty" xml:"DomainLevel,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The hostname.
	//
	// example:
	//
	// www
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMainDomainNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMainDomainNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetMainDomainNameResponseBody) GetDomainLevel() *int64 {
	return s.DomainLevel
}

func (s *GetMainDomainNameResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *GetMainDomainNameResponseBody) GetRR() *string {
	return s.RR
}

func (s *GetMainDomainNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMainDomainNameResponseBody) SetDomainLevel(v int64) *GetMainDomainNameResponseBody {
	s.DomainLevel = &v
	return s
}

func (s *GetMainDomainNameResponseBody) SetDomainName(v string) *GetMainDomainNameResponseBody {
	s.DomainName = &v
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
