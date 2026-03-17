// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickOutClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *KickOutClientsResponseBody
	GetRequestId() *string
}

type KickOutClientsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 76FD7E08-6AA1-4B1B-99FB-8B3CA6C99A8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KickOutClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KickOutClientsResponseBody) GoString() string {
	return s.String()
}

func (s *KickOutClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KickOutClientsResponseBody) SetRequestId(v string) *KickOutClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *KickOutClientsResponseBody) Validate() error {
	return dara.Validate(s)
}
