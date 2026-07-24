// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlFileResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlFileResponseBody) *GetSqlFileResponse
	GetBody() *GetSqlFileResponseBody
}

type GetSqlFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlFileResponse) GoString() string {
	return s.String()
}

func (s *GetSqlFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlFileResponse) GetBody() *GetSqlFileResponseBody {
	return s.Body
}

func (s *GetSqlFileResponse) SetHeaders(v map[string]*string) *GetSqlFileResponse {
	s.Headers = v
	return s
}

func (s *GetSqlFileResponse) SetStatusCode(v int32) *GetSqlFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlFileResponse) SetBody(v *GetSqlFileResponseBody) *GetSqlFileResponse {
	s.Body = v
	return s
}

func (s *GetSqlFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
