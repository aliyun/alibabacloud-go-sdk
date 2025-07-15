// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSchemaResponse
	GetStatusCode() *int32
	SetBody(v *CreateSchemaResponseBody) *CreateSchemaResponse
	GetBody() *CreateSchemaResponseBody
}

type CreateSchemaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemaResponse) GoString() string {
	return s.String()
}

func (s *CreateSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSchemaResponse) GetBody() *CreateSchemaResponseBody {
	return s.Body
}

func (s *CreateSchemaResponse) SetHeaders(v map[string]*string) *CreateSchemaResponse {
	s.Headers = v
	return s
}

func (s *CreateSchemaResponse) SetStatusCode(v int32) *CreateSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchemaResponse) SetBody(v *CreateSchemaResponseBody) *CreateSchemaResponse {
	s.Body = v
	return s
}

func (s *CreateSchemaResponse) Validate() error {
	return dara.Validate(s)
}
