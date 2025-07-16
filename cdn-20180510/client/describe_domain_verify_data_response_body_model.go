// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainVerifyDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeDomainVerifyDataResponseBody
	GetContent() *string
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
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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

func (s *DescribeDomainVerifyDataResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeDomainVerifyDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainVerifyDataResponseBody) SetContent(v string) *DescribeDomainVerifyDataResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDomainVerifyDataResponseBody) SetRequestId(v string) *DescribeDomainVerifyDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainVerifyDataResponseBody) Validate() error {
	return dara.Validate(s)
}
