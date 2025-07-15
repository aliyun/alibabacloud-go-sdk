// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSchemaPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSchemaPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSchemaPropertyResponse
	GetStatusCode() *int32
	SetBody(v *AddSchemaPropertyResponseBody) *AddSchemaPropertyResponse
	GetBody() *AddSchemaPropertyResponseBody
}

type AddSchemaPropertyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSchemaPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSchemaPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSchemaPropertyResponse) GoString() string {
	return s.String()
}

func (s *AddSchemaPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSchemaPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSchemaPropertyResponse) GetBody() *AddSchemaPropertyResponseBody {
	return s.Body
}

func (s *AddSchemaPropertyResponse) SetHeaders(v map[string]*string) *AddSchemaPropertyResponse {
	s.Headers = v
	return s
}

func (s *AddSchemaPropertyResponse) SetStatusCode(v int32) *AddSchemaPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSchemaPropertyResponse) SetBody(v *AddSchemaPropertyResponseBody) *AddSchemaPropertyResponse {
	s.Body = v
	return s
}

func (s *AddSchemaPropertyResponse) Validate() error {
	return dara.Validate(s)
}
