// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddDomainResponseBody
	GetDomainName() *string
	SetRequestId(v string) *AddDomainResponseBody
	GetRequestId() *string
}

type AddDomainResponseBody struct {
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// ADA27798-6911-4B06-AF34-53F62F62XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDomainResponseBody) SetDomainName(v string) *AddDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *AddDomainResponseBody) SetRequestId(v string) *AddDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
