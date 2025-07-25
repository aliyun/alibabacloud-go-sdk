// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsCacheDomainRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDnsCacheDomainRemarkResponseBody
	GetRequestId() *string
}

type UpdateDnsCacheDomainRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDnsCacheDomainRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsCacheDomainRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDnsCacheDomainRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDnsCacheDomainRemarkResponseBody) SetRequestId(v string) *UpdateDnsCacheDomainRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDnsCacheDomainRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
