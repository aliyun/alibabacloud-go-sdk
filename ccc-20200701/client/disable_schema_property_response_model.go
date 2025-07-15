// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSchemaPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSchemaPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSchemaPropertyResponse
	GetStatusCode() *int32
	SetBody(v *DisableSchemaPropertyResponseBody) *DisableSchemaPropertyResponse
	GetBody() *DisableSchemaPropertyResponseBody
}

type DisableSchemaPropertyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSchemaPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSchemaPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSchemaPropertyResponse) GoString() string {
	return s.String()
}

func (s *DisableSchemaPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSchemaPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSchemaPropertyResponse) GetBody() *DisableSchemaPropertyResponseBody {
	return s.Body
}

func (s *DisableSchemaPropertyResponse) SetHeaders(v map[string]*string) *DisableSchemaPropertyResponse {
	s.Headers = v
	return s
}

func (s *DisableSchemaPropertyResponse) SetStatusCode(v int32) *DisableSchemaPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSchemaPropertyResponse) SetBody(v *DisableSchemaPropertyResponseBody) *DisableSchemaPropertyResponse {
	s.Body = v
	return s
}

func (s *DisableSchemaPropertyResponse) Validate() error {
	return dara.Validate(s)
}
