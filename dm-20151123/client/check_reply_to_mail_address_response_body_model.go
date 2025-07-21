// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckReplyToMailAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckReplyToMailAddressResponseBody
	GetRequestId() *string
}

type CheckReplyToMailAddressResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckReplyToMailAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckReplyToMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CheckReplyToMailAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckReplyToMailAddressResponseBody) SetRequestId(v string) *CheckReplyToMailAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckReplyToMailAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
