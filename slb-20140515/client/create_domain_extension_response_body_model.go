// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainExtensionId(v string) *CreateDomainExtensionResponseBody
	GetDomainExtensionId() *string
	SetListenerPort(v int32) *CreateDomainExtensionResponseBody
	GetListenerPort() *int32
	SetRequestId(v string) *CreateDomainExtensionResponseBody
	GetRequestId() *string
}

type CreateDomainExtensionResponseBody struct {
	// The ID of the additional domain name.
	//
	// example:
	//
	// de-bp1rp7ta19******
	DomainExtensionId *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	// The frontend port that is used by the SLB instance.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A6E7EFC9-0938-40CA-877D-9BEDBD21D357
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainExtensionResponseBody) GetDomainExtensionId() *string {
	return s.DomainExtensionId
}

func (s *CreateDomainExtensionResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateDomainExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDomainExtensionResponseBody) SetDomainExtensionId(v string) *CreateDomainExtensionResponseBody {
	s.DomainExtensionId = &v
	return s
}

func (s *CreateDomainExtensionResponseBody) SetListenerPort(v int32) *CreateDomainExtensionResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *CreateDomainExtensionResponseBody) SetRequestId(v string) *CreateDomainExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDomainExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}
