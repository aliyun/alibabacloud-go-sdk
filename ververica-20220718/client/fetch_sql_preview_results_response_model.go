// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchSqlPreviewResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchSqlPreviewResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchSqlPreviewResultsResponse
	GetStatusCode() *int32
	SetBody(v *FetchSqlPreviewResultsResponseBody) *FetchSqlPreviewResultsResponse
	GetBody() *FetchSqlPreviewResultsResponseBody
}

type FetchSqlPreviewResultsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchSqlPreviewResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchSqlPreviewResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchSqlPreviewResultsResponse) GoString() string {
	return s.String()
}

func (s *FetchSqlPreviewResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchSqlPreviewResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchSqlPreviewResultsResponse) GetBody() *FetchSqlPreviewResultsResponseBody {
	return s.Body
}

func (s *FetchSqlPreviewResultsResponse) SetHeaders(v map[string]*string) *FetchSqlPreviewResultsResponse {
	s.Headers = v
	return s
}

func (s *FetchSqlPreviewResultsResponse) SetStatusCode(v int32) *FetchSqlPreviewResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchSqlPreviewResultsResponse) SetBody(v *FetchSqlPreviewResultsResponseBody) *FetchSqlPreviewResultsResponse {
	s.Body = v
	return s
}

func (s *FetchSqlPreviewResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
