// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateChatbotInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateChatbotInstanceResponseBody
	GetRequestId() *string
}

type AssociateChatbotInstanceResponseBody struct {
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateChatbotInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateChatbotInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateChatbotInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateChatbotInstanceResponseBody) SetRequestId(v string) *AssociateChatbotInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateChatbotInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
