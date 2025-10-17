// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BankMetaVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BankMetaVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *BankMetaVerifyIntlResponseBody) *BankMetaVerifyIntlResponse
	GetBody() *BankMetaVerifyIntlResponseBody
}

type BankMetaVerifyIntlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BankMetaVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BankMetaVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BankMetaVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BankMetaVerifyIntlResponse) GetBody() *BankMetaVerifyIntlResponseBody {
	return s.Body
}

func (s *BankMetaVerifyIntlResponse) SetHeaders(v map[string]*string) *BankMetaVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *BankMetaVerifyIntlResponse) SetStatusCode(v int32) *BankMetaVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *BankMetaVerifyIntlResponse) SetBody(v *BankMetaVerifyIntlResponseBody) *BankMetaVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *BankMetaVerifyIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
