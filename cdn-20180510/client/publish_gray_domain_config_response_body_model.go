// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishGrayDomainConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishGrayDomainConfigResponseBody
	GetRequestId() *string
}

type PublishGrayDomainConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishGrayDomainConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishGrayDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PublishGrayDomainConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishGrayDomainConfigResponseBody) SetRequestId(v string) *PublishGrayDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishGrayDomainConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
