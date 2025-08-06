// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultPlayDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultPlayDomainResponseBody
	GetRequestId() *string
}

type SetDefaultPlayDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultPlayDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultPlayDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultPlayDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultPlayDomainResponseBody) SetRequestId(v string) *SetDefaultPlayDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultPlayDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
