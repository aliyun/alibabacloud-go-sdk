// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopVodDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopVodDomainResponseBody
	GetRequestId() *string
}

type StopVodDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopVodDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopVodDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopVodDomainResponseBody) SetRequestId(v string) *StopVodDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopVodDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
