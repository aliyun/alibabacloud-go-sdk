// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemaPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSchemaPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSchemaPropertyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSchemaPropertyResponseBody) *DeleteSchemaPropertyResponse
	GetBody() *DeleteSchemaPropertyResponseBody
}

type DeleteSchemaPropertyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchemaPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchemaPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemaPropertyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchemaPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSchemaPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSchemaPropertyResponse) GetBody() *DeleteSchemaPropertyResponseBody {
	return s.Body
}

func (s *DeleteSchemaPropertyResponse) SetHeaders(v map[string]*string) *DeleteSchemaPropertyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchemaPropertyResponse) SetStatusCode(v int32) *DeleteSchemaPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchemaPropertyResponse) SetBody(v *DeleteSchemaPropertyResponseBody) *DeleteSchemaPropertyResponse {
	s.Body = v
	return s
}

func (s *DeleteSchemaPropertyResponse) Validate() error {
	return dara.Validate(s)
}
