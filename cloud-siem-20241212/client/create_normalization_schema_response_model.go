// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNormalizationSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNormalizationSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNormalizationSchemaResponse
	GetStatusCode() *int32
	SetBody(v *CreateNormalizationSchemaResponseBody) *CreateNormalizationSchemaResponse
	GetBody() *CreateNormalizationSchemaResponseBody
}

type CreateNormalizationSchemaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNormalizationSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNormalizationSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationSchemaResponse) GoString() string {
	return s.String()
}

func (s *CreateNormalizationSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNormalizationSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNormalizationSchemaResponse) GetBody() *CreateNormalizationSchemaResponseBody {
	return s.Body
}

func (s *CreateNormalizationSchemaResponse) SetHeaders(v map[string]*string) *CreateNormalizationSchemaResponse {
	s.Headers = v
	return s
}

func (s *CreateNormalizationSchemaResponse) SetStatusCode(v int32) *CreateNormalizationSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNormalizationSchemaResponse) SetBody(v *CreateNormalizationSchemaResponseBody) *CreateNormalizationSchemaResponse {
	s.Body = v
	return s
}

func (s *CreateNormalizationSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
