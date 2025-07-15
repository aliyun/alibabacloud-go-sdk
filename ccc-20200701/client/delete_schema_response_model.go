// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSchemaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSchemaResponseBody) *DeleteSchemaResponse
	GetBody() *DeleteSchemaResponseBody
}

type DeleteSchemaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemaResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSchemaResponse) GetBody() *DeleteSchemaResponseBody {
	return s.Body
}

func (s *DeleteSchemaResponse) SetHeaders(v map[string]*string) *DeleteSchemaResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchemaResponse) SetStatusCode(v int32) *DeleteSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchemaResponse) SetBody(v *DeleteSchemaResponseBody) *DeleteSchemaResponse {
	s.Body = v
	return s
}

func (s *DeleteSchemaResponse) Validate() error {
	return dara.Validate(s)
}
