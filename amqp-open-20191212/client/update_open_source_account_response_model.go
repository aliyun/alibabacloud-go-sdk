// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpenSourceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOpenSourceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOpenSourceAccountResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOpenSourceAccountResponseBody) *UpdateOpenSourceAccountResponse
	GetBody() *UpdateOpenSourceAccountResponseBody
}

type UpdateOpenSourceAccountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOpenSourceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOpenSourceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpenSourceAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateOpenSourceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOpenSourceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOpenSourceAccountResponse) GetBody() *UpdateOpenSourceAccountResponseBody {
	return s.Body
}

func (s *UpdateOpenSourceAccountResponse) SetHeaders(v map[string]*string) *UpdateOpenSourceAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateOpenSourceAccountResponse) SetStatusCode(v int32) *UpdateOpenSourceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOpenSourceAccountResponse) SetBody(v *UpdateOpenSourceAccountResponseBody) *UpdateOpenSourceAccountResponse {
	s.Body = v
	return s
}

func (s *UpdateOpenSourceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
