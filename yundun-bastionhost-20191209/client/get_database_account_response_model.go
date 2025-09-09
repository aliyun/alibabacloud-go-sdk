// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatabaseAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatabaseAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetDatabaseAccountResponseBody) *GetDatabaseAccountResponse
	GetBody() *GetDatabaseAccountResponseBody
}

type GetDatabaseAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatabaseAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabaseAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseAccountResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatabaseAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatabaseAccountResponse) GetBody() *GetDatabaseAccountResponseBody {
	return s.Body
}

func (s *GetDatabaseAccountResponse) SetHeaders(v map[string]*string) *GetDatabaseAccountResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseAccountResponse) SetStatusCode(v int32) *GetDatabaseAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseAccountResponse) SetBody(v *GetDatabaseAccountResponseBody) *GetDatabaseAccountResponse {
	s.Body = v
	return s
}

func (s *GetDatabaseAccountResponse) Validate() error {
	return dara.Validate(s)
}
