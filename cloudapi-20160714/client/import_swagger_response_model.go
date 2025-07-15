// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportSwaggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportSwaggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportSwaggerResponse
	GetStatusCode() *int32
	SetBody(v *ImportSwaggerResponseBody) *ImportSwaggerResponse
	GetBody() *ImportSwaggerResponseBody
}

type ImportSwaggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportSwaggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportSwaggerResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponse) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportSwaggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportSwaggerResponse) GetBody() *ImportSwaggerResponseBody {
	return s.Body
}

func (s *ImportSwaggerResponse) SetHeaders(v map[string]*string) *ImportSwaggerResponse {
	s.Headers = v
	return s
}

func (s *ImportSwaggerResponse) SetStatusCode(v int32) *ImportSwaggerResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportSwaggerResponse) SetBody(v *ImportSwaggerResponseBody) *ImportSwaggerResponse {
	s.Body = v
	return s
}

func (s *ImportSwaggerResponse) Validate() error {
	return dara.Validate(s)
}
