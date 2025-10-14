// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNormalizationCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNormalizationCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListNormalizationCategoriesResponseBody) *ListNormalizationCategoriesResponse
	GetBody() *ListNormalizationCategoriesResponseBody
}

type ListNormalizationCategoriesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNormalizationCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNormalizationCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListNormalizationCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNormalizationCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNormalizationCategoriesResponse) GetBody() *ListNormalizationCategoriesResponseBody {
	return s.Body
}

func (s *ListNormalizationCategoriesResponse) SetHeaders(v map[string]*string) *ListNormalizationCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListNormalizationCategoriesResponse) SetStatusCode(v int32) *ListNormalizationCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNormalizationCategoriesResponse) SetBody(v *ListNormalizationCategoriesResponseBody) *ListNormalizationCategoriesResponse {
	s.Body = v
	return s
}

func (s *ListNormalizationCategoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
