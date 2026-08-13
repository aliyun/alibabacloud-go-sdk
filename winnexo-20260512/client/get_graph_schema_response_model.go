// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGraphSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGraphSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetGraphSchemaResponseBody) *GetGraphSchemaResponse
	GetBody() *GetGraphSchemaResponseBody
}

type GetGraphSchemaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGraphSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGraphSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGraphSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGraphSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGraphSchemaResponse) GetBody() *GetGraphSchemaResponseBody {
	return s.Body
}

func (s *GetGraphSchemaResponse) SetHeaders(v map[string]*string) *GetGraphSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetGraphSchemaResponse) SetStatusCode(v int32) *GetGraphSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGraphSchemaResponse) SetBody(v *GetGraphSchemaResponseBody) *GetGraphSchemaResponse {
	s.Body = v
	return s
}

func (s *GetGraphSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
