// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSqlFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSqlFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSqlFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSqlFileResponseBody) *DeleteSqlFileResponse
	GetBody() *DeleteSqlFileResponseBody
}

type DeleteSqlFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSqlFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSqlFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSqlFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteSqlFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSqlFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSqlFileResponse) GetBody() *DeleteSqlFileResponseBody {
	return s.Body
}

func (s *DeleteSqlFileResponse) SetHeaders(v map[string]*string) *DeleteSqlFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteSqlFileResponse) SetStatusCode(v int32) *DeleteSqlFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSqlFileResponse) SetBody(v *DeleteSqlFileResponseBody) *DeleteSqlFileResponse {
	s.Body = v
	return s
}

func (s *DeleteSqlFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
