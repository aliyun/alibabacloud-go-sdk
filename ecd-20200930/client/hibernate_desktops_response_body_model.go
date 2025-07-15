// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHibernateDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *HibernateDesktopsResponseBody
	GetRequestId() *string
}

type HibernateDesktopsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 101AE027-8299-5E6E-A782-6C91C962****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HibernateDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HibernateDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *HibernateDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HibernateDesktopsResponseBody) SetRequestId(v string) *HibernateDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *HibernateDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
