// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatabaseAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatabaseAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatabaseAccountResponseBody) *DeleteDatabaseAccountResponse
	GetBody() *DeleteDatabaseAccountResponseBody
}

type DeleteDatabaseAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatabaseAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatabaseAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatabaseAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatabaseAccountResponse) GetBody() *DeleteDatabaseAccountResponseBody {
	return s.Body
}

func (s *DeleteDatabaseAccountResponse) SetHeaders(v map[string]*string) *DeleteDatabaseAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseAccountResponse) SetStatusCode(v int32) *DeleteDatabaseAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabaseAccountResponse) SetBody(v *DeleteDatabaseAccountResponseBody) *DeleteDatabaseAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteDatabaseAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
