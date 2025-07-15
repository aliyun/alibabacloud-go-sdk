// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSchemaPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSchemaPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSchemaPropertyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSchemaPropertyResponseBody) *UpdateSchemaPropertyResponse
	GetBody() *UpdateSchemaPropertyResponseBody
}

type UpdateSchemaPropertyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSchemaPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSchemaPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemaPropertyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSchemaPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSchemaPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSchemaPropertyResponse) GetBody() *UpdateSchemaPropertyResponseBody {
	return s.Body
}

func (s *UpdateSchemaPropertyResponse) SetHeaders(v map[string]*string) *UpdateSchemaPropertyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSchemaPropertyResponse) SetStatusCode(v int32) *UpdateSchemaPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSchemaPropertyResponse) SetBody(v *UpdateSchemaPropertyResponseBody) *UpdateSchemaPropertyResponse {
	s.Body = v
	return s
}

func (s *UpdateSchemaPropertyResponse) Validate() error {
	return dara.Validate(s)
}
