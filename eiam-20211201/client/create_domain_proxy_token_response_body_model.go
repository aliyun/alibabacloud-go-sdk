// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainProxyTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainProxyTokenId(v string) *CreateDomainProxyTokenResponseBody
	GetDomainProxyTokenId() *string
	SetRequestId(v string) *CreateDomainProxyTokenResponseBody
	GetRequestId() *string
}

type CreateDomainProxyTokenResponseBody struct {
	// The ID of the proxy token of the domain name.
	//
	// example:
	//
	// pt_mtohn73423stghoivjmi4jwxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainProxyTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainProxyTokenResponseBody) GetDomainProxyTokenId() *string {
	return s.DomainProxyTokenId
}

func (s *CreateDomainProxyTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDomainProxyTokenResponseBody) SetDomainProxyTokenId(v string) *CreateDomainProxyTokenResponseBody {
	s.DomainProxyTokenId = &v
	return s
}

func (s *CreateDomainProxyTokenResponseBody) SetRequestId(v string) *CreateDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDomainProxyTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
