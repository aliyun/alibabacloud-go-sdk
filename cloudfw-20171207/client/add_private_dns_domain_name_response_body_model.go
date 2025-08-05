// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivateDnsDomainNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddPrivateDnsDomainNameResponseBody
	GetRequestId() *string
}

type AddPrivateDnsDomainNameResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPrivateDnsDomainNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPrivateDnsDomainNameResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrivateDnsDomainNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPrivateDnsDomainNameResponseBody) SetRequestId(v string) *AddPrivateDnsDomainNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPrivateDnsDomainNameResponseBody) Validate() error {
	return dara.Validate(s)
}
