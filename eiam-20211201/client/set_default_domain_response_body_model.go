// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultDomainResponseBody
	GetRequestId() *string
}

type SetDefaultDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultDomainResponseBody) SetRequestId(v string) *SetDefaultDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
