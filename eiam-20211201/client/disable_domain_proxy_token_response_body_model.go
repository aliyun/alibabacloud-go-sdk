// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDomainProxyTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDomainProxyTokenResponseBody
	GetRequestId() *string
}

type DisableDomainProxyTokenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDomainProxyTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDomainProxyTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDomainProxyTokenResponseBody) SetRequestId(v string) *DisableDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDomainProxyTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
