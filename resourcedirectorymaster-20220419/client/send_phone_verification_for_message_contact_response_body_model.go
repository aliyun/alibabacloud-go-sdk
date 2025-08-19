// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPhoneVerificationForMessageContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendPhoneVerificationForMessageContactResponseBody
	GetRequestId() *string
}

type SendPhoneVerificationForMessageContactResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendPhoneVerificationForMessageContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendPhoneVerificationForMessageContactResponseBody) GoString() string {
	return s.String()
}

func (s *SendPhoneVerificationForMessageContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendPhoneVerificationForMessageContactResponseBody) SetRequestId(v string) *SendPhoneVerificationForMessageContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendPhoneVerificationForMessageContactResponseBody) Validate() error {
	return dara.Validate(s)
}
