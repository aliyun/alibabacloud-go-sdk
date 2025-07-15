// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTranscodeParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveDomainTranscodeParamsResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainTranscodeParamsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainTranscodeParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTranscodeParamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTranscodeParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainTranscodeParamsResponseBody) SetRequestId(v string) *DescribeLiveDomainTranscodeParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainTranscodeParamsResponseBody) Validate() error {
	return dara.Validate(s)
}
