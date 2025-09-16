// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlInstanceResponseBody) *GetSqlInstanceResponse
	GetBody() *GetSqlInstanceResponseBody
}

type GetSqlInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetSqlInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlInstanceResponse) GetBody() *GetSqlInstanceResponseBody {
	return s.Body
}

func (s *GetSqlInstanceResponse) SetHeaders(v map[string]*string) *GetSqlInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetSqlInstanceResponse) SetStatusCode(v int32) *GetSqlInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlInstanceResponse) SetBody(v *GetSqlInstanceResponseBody) *GetSqlInstanceResponse {
	s.Body = v
	return s
}

func (s *GetSqlInstanceResponse) Validate() error {
	return dara.Validate(s)
}
