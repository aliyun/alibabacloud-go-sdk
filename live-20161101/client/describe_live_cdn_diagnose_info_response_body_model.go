// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCdnDiagnoseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveCdnDiagnoseInfoResponseBody
	GetRequestId() *string
}

type DescribeLiveCdnDiagnoseInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveCdnDiagnoseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCdnDiagnoseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveCdnDiagnoseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveCdnDiagnoseInfoResponseBody) SetRequestId(v string) *DescribeLiveCdnDiagnoseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
