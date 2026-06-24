// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSourceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpenSourceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpenSourceAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpenSourceAccountResponseBody) *CreateOpenSourceAccountResponse
	GetBody() *CreateOpenSourceAccountResponseBody
}

type CreateOpenSourceAccountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpenSourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpenSourceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSourceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateOpenSourceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpenSourceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpenSourceAccountResponse) GetBody() *CreateOpenSourceAccountResponseBody {
	return s.Body
}

func (s *CreateOpenSourceAccountResponse) SetHeaders(v map[string]*string) *CreateOpenSourceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateOpenSourceAccountResponse) SetStatusCode(v int32) *CreateOpenSourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpenSourceAccountResponse) SetBody(v *CreateOpenSourceAccountResponseBody) *CreateOpenSourceAccountResponse {
	s.Body = v
	return s
}

func (s *CreateOpenSourceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
