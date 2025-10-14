// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNormalizationSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNormalizationSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNormalizationSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetNormalizationSchemaResponseBody) *GetNormalizationSchemaResponse
	GetBody() *GetNormalizationSchemaResponseBody
}

type GetNormalizationSchemaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNormalizationSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNormalizationSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNormalizationSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetNormalizationSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNormalizationSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNormalizationSchemaResponse) GetBody() *GetNormalizationSchemaResponseBody {
	return s.Body
}

func (s *GetNormalizationSchemaResponse) SetHeaders(v map[string]*string) *GetNormalizationSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetNormalizationSchemaResponse) SetStatusCode(v int32) *GetNormalizationSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNormalizationSchemaResponse) SetBody(v *GetNormalizationSchemaResponseBody) *GetNormalizationSchemaResponse {
	s.Body = v
	return s
}

func (s *GetNormalizationSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
