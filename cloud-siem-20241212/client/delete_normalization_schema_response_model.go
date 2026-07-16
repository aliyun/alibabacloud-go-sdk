// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNormalizationSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNormalizationSchemaResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNormalizationSchemaResponseBody) *DeleteNormalizationSchemaResponse
	GetBody() *DeleteNormalizationSchemaResponseBody
}

type DeleteNormalizationSchemaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNormalizationSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNormalizationSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationSchemaResponse) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNormalizationSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNormalizationSchemaResponse) GetBody() *DeleteNormalizationSchemaResponseBody {
	return s.Body
}

func (s *DeleteNormalizationSchemaResponse) SetHeaders(v map[string]*string) *DeleteNormalizationSchemaResponse {
	s.Headers = v
	return s
}

func (s *DeleteNormalizationSchemaResponse) SetStatusCode(v int32) *DeleteNormalizationSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNormalizationSchemaResponse) SetBody(v *DeleteNormalizationSchemaResponseBody) *DeleteNormalizationSchemaResponse {
	s.Body = v
	return s
}

func (s *DeleteNormalizationSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
