// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePrivateDNSResponseBody
	GetRequestId() *string
}

type UpdatePrivateDNSResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrivateDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateDNSResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrivateDNSResponseBody) SetRequestId(v string) *UpdatePrivateDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrivateDNSResponseBody) Validate() error {
	return dara.Validate(s)
}
