// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsCacheDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDnsCacheDomainResponseBody
	GetRequestId() *string
}

type DeleteDnsCacheDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDnsCacheDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsCacheDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDnsCacheDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDnsCacheDomainResponseBody) SetRequestId(v string) *DeleteDnsCacheDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDnsCacheDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
