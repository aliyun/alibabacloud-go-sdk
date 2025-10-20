// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDomainBrandResponseBody
	GetRequestId() *string
}

type UpdateDomainBrandResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDomainBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainBrandResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainBrandResponseBody) SetRequestId(v string) *UpdateDomainBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
