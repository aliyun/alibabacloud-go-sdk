// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomDomainSampleRateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomDomainSampleRateResponseBody
	GetRequestId() *string
}

type DeleteCustomDomainSampleRateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomDomainSampleRateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomDomainSampleRateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainSampleRateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomDomainSampleRateResponseBody) SetRequestId(v string) *DeleteCustomDomainSampleRateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomDomainSampleRateResponseBody) Validate() error {
	return dara.Validate(s)
}
