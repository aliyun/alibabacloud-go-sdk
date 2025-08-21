// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVerificationInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetVerificationInfoResponseBody
	GetRequestId() *string
}

type SetVerificationInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B182C041-8C64-5F2F-A07B-FC67FAF89CF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVerificationInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetVerificationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetVerificationInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetVerificationInfoResponseBody) SetRequestId(v string) *SetVerificationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetVerificationInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
