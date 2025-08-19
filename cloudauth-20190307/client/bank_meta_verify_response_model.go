// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BankMetaVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BankMetaVerifyResponse
	GetStatusCode() *int32
	SetBody(v *BankMetaVerifyResponseBody) *BankMetaVerifyResponse
	GetBody() *BankMetaVerifyResponseBody
}

type BankMetaVerifyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BankMetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BankMetaVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BankMetaVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BankMetaVerifyResponse) GetBody() *BankMetaVerifyResponseBody {
	return s.Body
}

func (s *BankMetaVerifyResponse) SetHeaders(v map[string]*string) *BankMetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *BankMetaVerifyResponse) SetStatusCode(v int32) *BankMetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *BankMetaVerifyResponse) SetBody(v *BankMetaVerifyResponseBody) *BankMetaVerifyResponse {
	s.Body = v
	return s
}

func (s *BankMetaVerifyResponse) Validate() error {
	return dara.Validate(s)
}
