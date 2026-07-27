// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKgSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportKgSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportKgSchemaResponse
	GetStatusCode() *int32
	SetBody(v *ImportKgSchemaResponseBody) *ImportKgSchemaResponse
	GetBody() *ImportKgSchemaResponseBody
}

type ImportKgSchemaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportKgSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportKgSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportKgSchemaResponse) GoString() string {
	return s.String()
}

func (s *ImportKgSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportKgSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportKgSchemaResponse) GetBody() *ImportKgSchemaResponseBody {
	return s.Body
}

func (s *ImportKgSchemaResponse) SetHeaders(v map[string]*string) *ImportKgSchemaResponse {
	s.Headers = v
	return s
}

func (s *ImportKgSchemaResponse) SetStatusCode(v int32) *ImportKgSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKgSchemaResponse) SetBody(v *ImportKgSchemaResponseBody) *ImportKgSchemaResponse {
	s.Body = v
	return s
}

func (s *ImportKgSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
