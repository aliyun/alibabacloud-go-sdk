// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNormalizationFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNormalizationFieldsResponse
	GetStatusCode() *int32
	SetBody(v *ListNormalizationFieldsResponseBody) *ListNormalizationFieldsResponse
	GetBody() *ListNormalizationFieldsResponseBody
}

type ListNormalizationFieldsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNormalizationFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNormalizationFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListNormalizationFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNormalizationFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNormalizationFieldsResponse) GetBody() *ListNormalizationFieldsResponseBody {
	return s.Body
}

func (s *ListNormalizationFieldsResponse) SetHeaders(v map[string]*string) *ListNormalizationFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListNormalizationFieldsResponse) SetStatusCode(v int32) *ListNormalizationFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNormalizationFieldsResponse) SetBody(v *ListNormalizationFieldsResponseBody) *ListNormalizationFieldsResponse {
	s.Body = v
	return s
}

func (s *ListNormalizationFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
