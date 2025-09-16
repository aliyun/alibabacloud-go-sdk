// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSqlInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSqlInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSqlInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSqlInstanceResponseBody) *DeleteSqlInstanceResponse
	GetBody() *DeleteSqlInstanceResponseBody
}

type DeleteSqlInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSqlInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSqlInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSqlInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSqlInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSqlInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSqlInstanceResponse) GetBody() *DeleteSqlInstanceResponseBody {
	return s.Body
}

func (s *DeleteSqlInstanceResponse) SetHeaders(v map[string]*string) *DeleteSqlInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSqlInstanceResponse) SetStatusCode(v int32) *DeleteSqlInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSqlInstanceResponse) SetBody(v *DeleteSqlInstanceResponseBody) *DeleteSqlInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteSqlInstanceResponse) Validate() error {
	return dara.Validate(s)
}
