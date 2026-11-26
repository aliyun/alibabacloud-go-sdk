// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendVerificationMessageResponseBody
	GetRequestId() *string
}

type SendVerificationMessageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A24A87F9-2B1C-5E4A-9D3E-1C6F8B902E41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendVerificationMessageResponseBody) SetRequestId(v string) *SendVerificationMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendVerificationMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
