// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSqlInstanceContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSqlInstanceContentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSqlInstanceContentResponseBody) *UpdateSqlInstanceContentResponse
	GetBody() *UpdateSqlInstanceContentResponseBody
}

type UpdateSqlInstanceContentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSqlInstanceContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSqlInstanceContentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSqlInstanceContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSqlInstanceContentResponse) GetBody() *UpdateSqlInstanceContentResponseBody {
	return s.Body
}

func (s *UpdateSqlInstanceContentResponse) SetHeaders(v map[string]*string) *UpdateSqlInstanceContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateSqlInstanceContentResponse) SetStatusCode(v int32) *UpdateSqlInstanceContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSqlInstanceContentResponse) SetBody(v *UpdateSqlInstanceContentResponseBody) *UpdateSqlInstanceContentResponse {
	s.Body = v
	return s
}

func (s *UpdateSqlInstanceContentResponse) Validate() error {
	return dara.Validate(s)
}
