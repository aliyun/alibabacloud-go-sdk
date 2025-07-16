// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationNoTableFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchFormDataSecondGenerationNoTableFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchFormDataSecondGenerationNoTableFieldResponse
	GetStatusCode() *int32
	SetBody(v *SearchFormDataSecondGenerationNoTableFieldResponseBody) *SearchFormDataSecondGenerationNoTableFieldResponse
	GetBody() *SearchFormDataSecondGenerationNoTableFieldResponseBody
}

type SearchFormDataSecondGenerationNoTableFieldResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDataSecondGenerationNoTableFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) GetBody() *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	return s.Body
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) SetHeaders(v map[string]*string) *SearchFormDataSecondGenerationNoTableFieldResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) SetStatusCode(v int32) *SearchFormDataSecondGenerationNoTableFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) SetBody(v *SearchFormDataSecondGenerationNoTableFieldResponseBody) *SearchFormDataSecondGenerationNoTableFieldResponse {
	s.Body = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) Validate() error {
	return dara.Validate(s)
}
