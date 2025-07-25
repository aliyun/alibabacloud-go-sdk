// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDomainRemarkResponseBody
	GetRequestId() *string
}

type UpdateDomainRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainRemarkResponseBody) SetRequestId(v string) *UpdateDomainRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
