// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJDBCDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteJDBCDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteJDBCDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteJDBCDataSourceResponseBody) *DeleteJDBCDataSourceResponse
	GetBody() *DeleteJDBCDataSourceResponseBody
}

type DeleteJDBCDataSourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJDBCDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJDBCDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteJDBCDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteJDBCDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteJDBCDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteJDBCDataSourceResponse) GetBody() *DeleteJDBCDataSourceResponseBody {
	return s.Body
}

func (s *DeleteJDBCDataSourceResponse) SetHeaders(v map[string]*string) *DeleteJDBCDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteJDBCDataSourceResponse) SetStatusCode(v int32) *DeleteJDBCDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJDBCDataSourceResponse) SetBody(v *DeleteJDBCDataSourceResponseBody) *DeleteJDBCDataSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteJDBCDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
