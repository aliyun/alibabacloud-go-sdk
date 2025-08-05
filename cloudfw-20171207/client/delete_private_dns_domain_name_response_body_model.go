// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsDomainNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrivateDnsDomainNameResponseBody
	GetRequestId() *string
}

type DeletePrivateDnsDomainNameResponseBody struct {
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-0009012****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateDnsDomainNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsDomainNameResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsDomainNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateDnsDomainNameResponseBody) SetRequestId(v string) *DeletePrivateDnsDomainNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateDnsDomainNameResponseBody) Validate() error {
	return dara.Validate(s)
}
