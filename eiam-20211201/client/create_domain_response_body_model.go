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
	// Domain ID.
	//
	// example:
	//
	// dm_mtohn6mltdz3ibtly2rxvnvxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
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
