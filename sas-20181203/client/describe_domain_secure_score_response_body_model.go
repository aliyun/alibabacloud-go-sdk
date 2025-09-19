// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureScoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainSecureScoreResponseBody
	GetRequestId() *string
	SetSecurityScore(v int32) *DescribeDomainSecureScoreResponseBody
	GetSecurityScore() *int32
}

type DescribeDomainSecureScoreResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 11C96623-E106-59C9-866D-A6C82911****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security score of the website.
	//
	// example:
	//
	// 100
	SecurityScore *int32 `json:"SecurityScore,omitempty" xml:"SecurityScore,omitempty"`
}

func (s DescribeDomainSecureScoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureScoreResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureScoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSecureScoreResponseBody) GetSecurityScore() *int32 {
	return s.SecurityScore
}

func (s *DescribeDomainSecureScoreResponseBody) SetRequestId(v string) *DescribeDomainSecureScoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSecureScoreResponseBody) SetSecurityScore(v int32) *DescribeDomainSecureScoreResponseBody {
	s.SecurityScore = &v
	return s
}

func (s *DescribeDomainSecureScoreResponseBody) Validate() error {
	return dara.Validate(s)
}
