// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSqlInstanceNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSqlInstanceNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSqlInstanceNameResponseBody) *UpdateSqlInstanceNameResponse
	GetBody() *UpdateSqlInstanceNameResponseBody
}

type UpdateSqlInstanceNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSqlInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSqlInstanceNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSqlInstanceNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSqlInstanceNameResponse) GetBody() *UpdateSqlInstanceNameResponseBody {
	return s.Body
}

func (s *UpdateSqlInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateSqlInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateSqlInstanceNameResponse) SetStatusCode(v int32) *UpdateSqlInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSqlInstanceNameResponse) SetBody(v *UpdateSqlInstanceNameResponseBody) *UpdateSqlInstanceNameResponse {
	s.Body = v
	return s
}

func (s *UpdateSqlInstanceNameResponse) Validate() error {
	return dara.Validate(s)
}
