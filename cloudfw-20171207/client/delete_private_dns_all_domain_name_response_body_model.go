// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsAllDomainNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrivateDnsAllDomainNameResponseBody
	GetRequestId() *string
}

type DeletePrivateDnsAllDomainNameResponseBody struct {
	// example:
	//
	// B2841452-CB8D-4F7D-B247-38E1CF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateDnsAllDomainNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsAllDomainNameResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsAllDomainNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateDnsAllDomainNameResponseBody) SetRequestId(v string) *DeletePrivateDnsAllDomainNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateDnsAllDomainNameResponseBody) Validate() error {
	return dara.Validate(s)
}
