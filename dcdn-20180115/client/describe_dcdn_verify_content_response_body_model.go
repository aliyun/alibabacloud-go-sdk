// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnVerifyContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DescribeDcdnVerifyContentResponseBody
	GetContent() *string
	SetRequestId(v string) *DescribeDcdnVerifyContentResponseBody
	GetRequestId() *string
}

type DescribeDcdnVerifyContentResponseBody struct {
	// The verification result.
	//
	// example:
	//
	// verify_dffeb6610035dcb78b413a59c31cd9**
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnVerifyContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnVerifyContentResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeDcdnVerifyContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnVerifyContentResponseBody) SetContent(v string) *DescribeDcdnVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeDcdnVerifyContentResponseBody) SetRequestId(v string) *DescribeDcdnVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnVerifyContentResponseBody) Validate() error {
	return dara.Validate(s)
}
