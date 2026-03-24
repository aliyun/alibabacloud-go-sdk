// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainVerifyDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeDomainVerifyDataResponseBodyContent) *DescribeDomainVerifyDataResponseBody
	GetContent() *DescribeDomainVerifyDataResponseBodyContent
	SetRequestId(v string) *DescribeDomainVerifyDataResponseBody
	GetRequestId() *string
}

type DescribeDomainVerifyDataResponseBody struct {
	// The verification content.
	//
	// example:
	//
	// {
	//
	//     "verifiCode": "uy0-DbxL4HBmUtSUXpkXctaSnCAUKhhNH6WKl-JnJY4",
	//
	//     "verifyKey": "_acme-challenge"
	//
	//   }
	Content *DescribeDomainVerifyDataResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainVerifyDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainVerifyDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainVerifyDataResponseBody) GetContent() *DescribeDomainVerifyDataResponseBodyContent {
	return s.Content
}

func (s *DescribeDomainVerifyDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainVerifyDataResponseBody) SetContent(v *DescribeDomainVerifyDataResponseBodyContent) *DescribeDomainVerifyDataResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDomainVerifyDataResponseBody) SetRequestId(v string) *DescribeDomainVerifyDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainVerifyDataResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainVerifyDataResponseBodyContent struct {
	RootDomain *string `json:"RootDomain,omitempty" xml:"RootDomain,omitempty"`
	VerifyCode *string `json:"verifyCode,omitempty" xml:"verifyCode,omitempty"`
	VerifyKey  *string `json:"verifyKey,omitempty" xml:"verifyKey,omitempty"`
}

func (s DescribeDomainVerifyDataResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainVerifyDataResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDomainVerifyDataResponseBodyContent) GetRootDomain() *string {
	return s.RootDomain
}

func (s *DescribeDomainVerifyDataResponseBodyContent) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *DescribeDomainVerifyDataResponseBodyContent) GetVerifyKey() *string {
	return s.VerifyKey
}

func (s *DescribeDomainVerifyDataResponseBodyContent) SetRootDomain(v string) *DescribeDomainVerifyDataResponseBodyContent {
	s.RootDomain = &v
	return s
}

func (s *DescribeDomainVerifyDataResponseBodyContent) SetVerifyCode(v string) *DescribeDomainVerifyDataResponseBodyContent {
	s.VerifyCode = &v
	return s
}

func (s *DescribeDomainVerifyDataResponseBodyContent) SetVerifyKey(v string) *DescribeDomainVerifyDataResponseBodyContent {
	s.VerifyKey = &v
	return s
}

func (s *DescribeDomainVerifyDataResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
