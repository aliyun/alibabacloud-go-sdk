// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainProxyTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDomainProxyTokenResponseBody
	GetRequestId() *string
}

type DeleteDomainProxyTokenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainProxyTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainProxyTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainProxyTokenResponseBody) SetRequestId(v string) *DeleteDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainProxyTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
