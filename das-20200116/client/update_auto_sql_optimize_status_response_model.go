// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoSqlOptimizeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoSqlOptimizeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoSqlOptimizeStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoSqlOptimizeStatusResponseBody) *UpdateAutoSqlOptimizeStatusResponse
	GetBody() *UpdateAutoSqlOptimizeStatusResponseBody
}

type UpdateAutoSqlOptimizeStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoSqlOptimizeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoSqlOptimizeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoSqlOptimizeStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoSqlOptimizeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoSqlOptimizeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoSqlOptimizeStatusResponse) GetBody() *UpdateAutoSqlOptimizeStatusResponseBody {
	return s.Body
}

func (s *UpdateAutoSqlOptimizeStatusResponse) SetHeaders(v map[string]*string) *UpdateAutoSqlOptimizeStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponse) SetStatusCode(v int32) *UpdateAutoSqlOptimizeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponse) SetBody(v *UpdateAutoSqlOptimizeStatusResponseBody) *UpdateAutoSqlOptimizeStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponse) Validate() error {
	return dara.Validate(s)
}
