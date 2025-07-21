// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *CreateDomainResponseBody
	GetDomainId() *string
	SetRequestId(v string) *CreateDomainResponseBody
	GetRequestId() *string
}

type CreateDomainResponseBody struct {
	// Domain ID
	//
	// example:
	//
	// 158910
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// Request ID
	//
	// example:
	//
	// B49AD828-25D1-488C-90B7-8853C1944486
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDomainResponseBody) SetDomainId(v string) *CreateDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
