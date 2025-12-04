// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSQLAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSQLAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateSQLAccountResponseBody) *CreateSQLAccountResponse
	GetBody() *CreateSQLAccountResponseBody
}

type CreateSQLAccountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSQLAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSQLAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateSQLAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSQLAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSQLAccountResponse) GetBody() *CreateSQLAccountResponseBody {
	return s.Body
}

func (s *CreateSQLAccountResponse) SetHeaders(v map[string]*string) *CreateSQLAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateSQLAccountResponse) SetStatusCode(v int32) *CreateSQLAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSQLAccountResponse) SetBody(v *CreateSQLAccountResponseBody) *CreateSQLAccountResponse {
	s.Body = v
	return s
}

func (s *CreateSQLAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
