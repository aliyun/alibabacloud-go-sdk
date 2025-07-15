// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetSchemaResponseBody) *GetSchemaResponse
	GetBody() *GetSchemaResponseBody
}

type GetSchemaResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSchemaResponse) GetBody() *GetSchemaResponseBody {
	return s.Body
}

func (s *GetSchemaResponse) SetHeaders(v map[string]*string) *GetSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetSchemaResponse) SetStatusCode(v int32) *GetSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSchemaResponse) SetBody(v *GetSchemaResponseBody) *GetSchemaResponse {
	s.Body = v
	return s
}

func (s *GetSchemaResponse) Validate() error {
	return dara.Validate(s)
}
