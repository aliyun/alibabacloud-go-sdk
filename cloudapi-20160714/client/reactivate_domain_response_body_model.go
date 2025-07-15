// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReactivateDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReactivateDomainResponseBody
	GetRequestId() *string
}

type ReactivateDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReactivateDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReactivateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ReactivateDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReactivateDomainResponseBody) SetRequestId(v string) *ReactivateDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReactivateDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
