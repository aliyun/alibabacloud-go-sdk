// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSqlInstanceParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSqlInstanceParamsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSqlInstanceParamsResponseBody) *UpdateSqlInstanceParamsResponse
	GetBody() *UpdateSqlInstanceParamsResponseBody
}

type UpdateSqlInstanceParamsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSqlInstanceParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSqlInstanceParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceParamsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSqlInstanceParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSqlInstanceParamsResponse) GetBody() *UpdateSqlInstanceParamsResponseBody {
	return s.Body
}

func (s *UpdateSqlInstanceParamsResponse) SetHeaders(v map[string]*string) *UpdateSqlInstanceParamsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSqlInstanceParamsResponse) SetStatusCode(v int32) *UpdateSqlInstanceParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSqlInstanceParamsResponse) SetBody(v *UpdateSqlInstanceParamsResponseBody) *UpdateSqlInstanceParamsResponse {
	s.Body = v
	return s
}

func (s *UpdateSqlInstanceParamsResponse) Validate() error {
	return dara.Validate(s)
}
