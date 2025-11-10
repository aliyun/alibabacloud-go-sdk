// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountResponseBody) *GetAccountResponse
	GetBody() *GetAccountResponseBody
}

type GetAccountResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountResponse) GoString() string {
	return s.String()
}

func (s *GetAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountResponse) GetBody() *GetAccountResponseBody {
	return s.Body
}

func (s *GetAccountResponse) SetHeaders(v map[string]*string) *GetAccountResponse {
	s.Headers = v
	return s
}

func (s *GetAccountResponse) SetStatusCode(v int32) *GetAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountResponse) SetBody(v *GetAccountResponseBody) *GetAccountResponse {
	s.Body = v
	return s
}

func (s *GetAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
