// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationCodeForEnableRDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendVerificationCodeForEnableRDResponseBody
	GetRequestId() *string
}

type SendVerificationCodeForEnableRDResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC2FE94D-A4A2-51A1-A493-5C273A36C46A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationCodeForEnableRDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationCodeForEnableRDResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForEnableRDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendVerificationCodeForEnableRDResponseBody) SetRequestId(v string) *SendVerificationCodeForEnableRDResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendVerificationCodeForEnableRDResponseBody) Validate() error {
	return dara.Validate(s)
}
