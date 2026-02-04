// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNormalizationSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNormalizationSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNormalizationSchemaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNormalizationSchemaResponseBody) *UpdateNormalizationSchemaResponse
	GetBody() *UpdateNormalizationSchemaResponseBody
}

type UpdateNormalizationSchemaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNormalizationSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNormalizationSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationSchemaResponse) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNormalizationSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNormalizationSchemaResponse) GetBody() *UpdateNormalizationSchemaResponseBody {
	return s.Body
}

func (s *UpdateNormalizationSchemaResponse) SetHeaders(v map[string]*string) *UpdateNormalizationSchemaResponse {
	s.Headers = v
	return s
}

func (s *UpdateNormalizationSchemaResponse) SetStatusCode(v int32) *UpdateNormalizationSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNormalizationSchemaResponse) SetBody(v *UpdateNormalizationSchemaResponseBody) *UpdateNormalizationSchemaResponse {
	s.Body = v
	return s
}

func (s *UpdateNormalizationSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
