// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendEmailVerificationForMessageContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendEmailVerificationForMessageContactResponseBody
	GetRequestId() *string
}

type SendEmailVerificationForMessageContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendEmailVerificationForMessageContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendEmailVerificationForMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *SendEmailVerificationForMessageContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendEmailVerificationForMessageContactResponseBody) SetRequestId(v string) *SendEmailVerificationForMessageContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendEmailVerificationForMessageContactResponseBody) Validate() error {
	return dara.Validate(s)
}
