// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityResultsResponseBody) *ListDataQualityResultsResponse
	GetBody() *ListDataQualityResultsResponseBody
}

type ListDataQualityResultsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityResultsResponse) GetBody() *ListDataQualityResultsResponseBody {
	return s.Body
}

func (s *ListDataQualityResultsResponse) SetHeaders(v map[string]*string) *ListDataQualityResultsResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityResultsResponse) SetStatusCode(v int32) *ListDataQualityResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityResultsResponse) SetBody(v *ListDataQualityResultsResponseBody) *ListDataQualityResultsResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
