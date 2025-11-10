// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePayerForAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePayerForAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePayerForAccountResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePayerForAccountResponseBody) *UpdatePayerForAccountResponse
	GetBody() *UpdatePayerForAccountResponseBody
}

type UpdatePayerForAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePayerForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePayerForAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePayerForAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdatePayerForAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePayerForAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePayerForAccountResponse) GetBody() *UpdatePayerForAccountResponseBody {
	return s.Body
}

func (s *UpdatePayerForAccountResponse) SetHeaders(v map[string]*string) *UpdatePayerForAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdatePayerForAccountResponse) SetStatusCode(v int32) *UpdatePayerForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePayerForAccountResponse) SetBody(v *UpdatePayerForAccountResponseBody) *UpdatePayerForAccountResponse {
	s.Body = v
	return s
}

func (s *UpdatePayerForAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
