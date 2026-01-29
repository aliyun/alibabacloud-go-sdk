// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgAccountResponseBody) *CreateAgAccountResponse
	GetBody() *CreateAgAccountResponseBody
}

type CreateAgAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgAccountResponse) GetBody() *CreateAgAccountResponseBody {
	return s.Body
}

func (s *CreateAgAccountResponse) SetHeaders(v map[string]*string) *CreateAgAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAgAccountResponse) SetStatusCode(v int32) *CreateAgAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgAccountResponse) SetBody(v *CreateAgAccountResponseBody) *CreateAgAccountResponse {
	s.Body = v
	return s
}

func (s *CreateAgAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
