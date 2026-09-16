// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLumaWithSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLumaWithSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLumaWithSQLResponse
	GetStatusCode() *int32
	SetBody(v *QueryLumaWithSQLResponseBody) *QueryLumaWithSQLResponse
	GetBody() *QueryLumaWithSQLResponseBody
}

type QueryLumaWithSQLResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLumaWithSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLumaWithSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLumaWithSQLResponse) GoString() string {
	return s.String()
}

func (s *QueryLumaWithSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLumaWithSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLumaWithSQLResponse) GetBody() *QueryLumaWithSQLResponseBody {
	return s.Body
}

func (s *QueryLumaWithSQLResponse) SetHeaders(v map[string]*string) *QueryLumaWithSQLResponse {
	s.Headers = v
	return s
}

func (s *QueryLumaWithSQLResponse) SetStatusCode(v int32) *QueryLumaWithSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLumaWithSQLResponse) SetBody(v *QueryLumaWithSQLResponseBody) *QueryLumaWithSQLResponse {
	s.Body = v
	return s
}

func (s *QueryLumaWithSQLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
