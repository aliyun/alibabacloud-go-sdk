// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatasetDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchDatasetDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchDatasetDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *SearchDatasetDocumentsResponseBody) *SearchDatasetDocumentsResponse
	GetBody() *SearchDatasetDocumentsResponseBody
}

type SearchDatasetDocumentsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDatasetDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDatasetDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchDatasetDocumentsResponse) GoString() string {
	return s.String()
}

func (s *SearchDatasetDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchDatasetDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchDatasetDocumentsResponse) GetBody() *SearchDatasetDocumentsResponseBody {
	return s.Body
}

func (s *SearchDatasetDocumentsResponse) SetHeaders(v map[string]*string) *SearchDatasetDocumentsResponse {
	s.Headers = v
	return s
}

func (s *SearchDatasetDocumentsResponse) SetStatusCode(v int32) *SearchDatasetDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDatasetDocumentsResponse) SetBody(v *SearchDatasetDocumentsResponseBody) *SearchDatasetDocumentsResponse {
	s.Body = v
	return s
}

func (s *SearchDatasetDocumentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
