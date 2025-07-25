// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsCacheDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDnsCacheDomainResponseBody
	GetRequestId() *string
}

type UpdateDnsCacheDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDnsCacheDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsCacheDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsCacheDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDnsCacheDomainResponseBody) SetRequestId(v string) *UpdateDnsCacheDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDnsCacheDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
