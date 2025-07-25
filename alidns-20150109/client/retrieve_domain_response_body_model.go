// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RetrieveDomainResponseBody
	GetRequestId() *string
}

type RetrieveDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9CC0D642-49D4-48DE-A1A5-9F218652E4A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetrieveDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveDomainResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrieveDomainResponseBody) SetRequestId(v string) *RetrieveDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
