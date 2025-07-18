// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJDBCDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJDBCDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJDBCDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateJDBCDataSourceResponseBody) *CreateJDBCDataSourceResponse
	GetBody() *CreateJDBCDataSourceResponseBody
}

type CreateJDBCDataSourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJDBCDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJDBCDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJDBCDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateJDBCDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJDBCDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJDBCDataSourceResponse) GetBody() *CreateJDBCDataSourceResponseBody {
	return s.Body
}

func (s *CreateJDBCDataSourceResponse) SetHeaders(v map[string]*string) *CreateJDBCDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateJDBCDataSourceResponse) SetStatusCode(v int32) *CreateJDBCDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJDBCDataSourceResponse) SetBody(v *CreateJDBCDataSourceResponseBody) *CreateJDBCDataSourceResponse {
	s.Body = v
	return s
}

func (s *CreateJDBCDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
