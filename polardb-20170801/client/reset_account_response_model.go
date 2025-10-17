// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAccountResponse
	GetStatusCode() *int32
	SetBody(v *ResetAccountResponseBody) *ResetAccountResponse
	GetBody() *ResetAccountResponseBody
}

type ResetAccountResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAccountResponse) GetBody() *ResetAccountResponseBody {
	return s.Body
}

func (s *ResetAccountResponse) SetHeaders(v map[string]*string) *ResetAccountResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountResponse) SetStatusCode(v int32) *ResetAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountResponse) SetBody(v *ResetAccountResponseBody) *ResetAccountResponse {
	s.Body = v
	return s
}

func (s *ResetAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
