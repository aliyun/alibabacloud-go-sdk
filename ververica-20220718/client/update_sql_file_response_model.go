// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSqlFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSqlFileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSqlFileResponseBody) *UpdateSqlFileResponse
	GetBody() *UpdateSqlFileResponseBody
}

type UpdateSqlFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSqlFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSqlFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateSqlFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSqlFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSqlFileResponse) GetBody() *UpdateSqlFileResponseBody {
	return s.Body
}

func (s *UpdateSqlFileResponse) SetHeaders(v map[string]*string) *UpdateSqlFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateSqlFileResponse) SetStatusCode(v int32) *UpdateSqlFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSqlFileResponse) SetBody(v *UpdateSqlFileResponseBody) *UpdateSqlFileResponse {
	s.Body = v
	return s
}

func (s *UpdateSqlFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
