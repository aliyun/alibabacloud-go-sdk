// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpenSearchAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpenSearchAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpenSearchAccountResponseBody) *CreateOpenSearchAccountResponse
	GetBody() *CreateOpenSearchAccountResponseBody
}

type CreateOpenSearchAccountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpenSearchAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpenSearchAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpenSearchAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpenSearchAccountResponse) GetBody() *CreateOpenSearchAccountResponseBody {
	return s.Body
}

func (s *CreateOpenSearchAccountResponse) SetHeaders(v map[string]*string) *CreateOpenSearchAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateOpenSearchAccountResponse) SetStatusCode(v int32) *CreateOpenSearchAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpenSearchAccountResponse) SetBody(v *CreateOpenSearchAccountResponseBody) *CreateOpenSearchAccountResponse {
	s.Body = v
	return s
}

func (s *CreateOpenSearchAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
