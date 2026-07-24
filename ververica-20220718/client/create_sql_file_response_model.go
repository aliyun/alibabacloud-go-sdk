// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSqlFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSqlFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateSqlFileResponseBody) *CreateSqlFileResponse
	GetBody() *CreateSqlFileResponseBody
}

type CreateSqlFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSqlFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSqlFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlFileResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSqlFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSqlFileResponse) GetBody() *CreateSqlFileResponseBody {
	return s.Body
}

func (s *CreateSqlFileResponse) SetHeaders(v map[string]*string) *CreateSqlFileResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlFileResponse) SetStatusCode(v int32) *CreateSqlFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSqlFileResponse) SetBody(v *CreateSqlFileResponseBody) *CreateSqlFileResponse {
	s.Body = v
	return s
}

func (s *CreateSqlFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
