// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainDnssecStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDomainDnssecStatusResponseBody
	GetRequestId() *string
}

type SetDomainDnssecStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainDnssecStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDomainDnssecStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainDnssecStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDomainDnssecStatusResponseBody) SetRequestId(v string) *SetDomainDnssecStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDomainDnssecStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
