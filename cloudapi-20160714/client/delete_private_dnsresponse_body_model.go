// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrivateDNSResponseBody
	GetRequestId() *string
}

type DeletePrivateDNSResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDNSResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateDNSResponseBody) SetRequestId(v string) *DeletePrivateDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateDNSResponseBody) Validate() error {
	return dara.Validate(s)
}
