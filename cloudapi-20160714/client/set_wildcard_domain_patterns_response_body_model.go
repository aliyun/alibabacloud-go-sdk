// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWildcardDomainPatternsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetWildcardDomainPatternsResponseBody
	GetRequestId() *string
}

type SetWildcardDomainPatternsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D0075BDA-8AED-5073-A70A-FE44E86AB20F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetWildcardDomainPatternsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetWildcardDomainPatternsResponseBody) GoString() string {
	return s.String()
}

func (s *SetWildcardDomainPatternsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetWildcardDomainPatternsResponseBody) SetRequestId(v string) *SetWildcardDomainPatternsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetWildcardDomainPatternsResponseBody) Validate() error {
	return dara.Validate(s)
}
