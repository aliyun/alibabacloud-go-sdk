// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMailAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMailAddressId(v string) *CreateMailAddressResponseBody
	GetMailAddressId() *string
	SetRequestId(v string) *CreateMailAddressResponseBody
	GetRequestId() *string
}

type CreateMailAddressResponseBody struct {
	// Mail address ID
	//
	// example:
	//
	// 15123
	MailAddressId *string `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 95A7D497-F8DD-4834-B81E-C1783236E55F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMailAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMailAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMailAddressResponseBody) GetMailAddressId() *string {
	return s.MailAddressId
}

func (s *CreateMailAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMailAddressResponseBody) SetMailAddressId(v string) *CreateMailAddressResponseBody {
	s.MailAddressId = &v
	return s
}

func (s *CreateMailAddressResponseBody) SetRequestId(v string) *CreateMailAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMailAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
