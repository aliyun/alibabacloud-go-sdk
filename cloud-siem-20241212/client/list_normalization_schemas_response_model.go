// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNormalizationSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNormalizationSchemasResponse
	GetStatusCode() *int32
	SetBody(v *ListNormalizationSchemasResponseBody) *ListNormalizationSchemasResponse
	GetBody() *ListNormalizationSchemasResponseBody
}

type ListNormalizationSchemasResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNormalizationSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNormalizationSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListNormalizationSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNormalizationSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNormalizationSchemasResponse) GetBody() *ListNormalizationSchemasResponseBody {
	return s.Body
}

func (s *ListNormalizationSchemasResponse) SetHeaders(v map[string]*string) *ListNormalizationSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListNormalizationSchemasResponse) SetStatusCode(v int32) *ListNormalizationSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNormalizationSchemasResponse) SetBody(v *ListNormalizationSchemasResponseBody) *ListNormalizationSchemasResponse {
	s.Body = v
	return s
}

func (s *ListNormalizationSchemasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
