// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartVodDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartVodDomainResponseBody
	GetRequestId() *string
}

type StartVodDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartVodDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartVodDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartVodDomainResponseBody) SetRequestId(v string) *StartVodDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartVodDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
