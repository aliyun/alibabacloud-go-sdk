// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchFormDataSecondGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchFormDataSecondGenerationResponse
	GetStatusCode() *int32
	SetBody(v *SearchFormDataSecondGenerationResponseBody) *SearchFormDataSecondGenerationResponse
	GetBody() *SearchFormDataSecondGenerationResponseBody
}

type SearchFormDataSecondGenerationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDataSecondGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDataSecondGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchFormDataSecondGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchFormDataSecondGenerationResponse) GetBody() *SearchFormDataSecondGenerationResponseBody {
	return s.Body
}

func (s *SearchFormDataSecondGenerationResponse) SetHeaders(v map[string]*string) *SearchFormDataSecondGenerationResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataSecondGenerationResponse) SetStatusCode(v int32) *SearchFormDataSecondGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponse) SetBody(v *SearchFormDataSecondGenerationResponseBody) *SearchFormDataSecondGenerationResponse {
	s.Body = v
	return s
}

func (s *SearchFormDataSecondGenerationResponse) Validate() error {
	return dara.Validate(s)
}
