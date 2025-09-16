// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSqlInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSqlInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateSqlInstanceResponseBody) *CreateSqlInstanceResponse
	GetBody() *CreateSqlInstanceResponseBody
}

type CreateSqlInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSqlInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSqlInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSqlInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSqlInstanceResponse) GetBody() *CreateSqlInstanceResponseBody {
	return s.Body
}

func (s *CreateSqlInstanceResponse) SetHeaders(v map[string]*string) *CreateSqlInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlInstanceResponse) SetStatusCode(v int32) *CreateSqlInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSqlInstanceResponse) SetBody(v *CreateSqlInstanceResponseBody) *CreateSqlInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateSqlInstanceResponse) Validate() error {
	return dara.Validate(s)
}
