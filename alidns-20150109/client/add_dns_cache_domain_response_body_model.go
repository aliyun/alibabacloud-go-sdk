// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsCacheDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDnsCacheDomainResponseBody
	GetRequestId() *string
}

type AddDnsCacheDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDnsCacheDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDnsCacheDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDnsCacheDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDnsCacheDomainResponseBody) SetRequestId(v string) *AddDnsCacheDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDnsCacheDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
