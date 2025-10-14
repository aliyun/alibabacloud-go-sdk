// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataIngestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataIngestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataIngestionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataIngestionsResponseBody) *ListDataIngestionsResponse
	GetBody() *ListDataIngestionsResponseBody
}

type ListDataIngestionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataIngestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataIngestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionsResponse) GoString() string {
	return s.String()
}

func (s *ListDataIngestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataIngestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataIngestionsResponse) GetBody() *ListDataIngestionsResponseBody {
	return s.Body
}

func (s *ListDataIngestionsResponse) SetHeaders(v map[string]*string) *ListDataIngestionsResponse {
	s.Headers = v
	return s
}

func (s *ListDataIngestionsResponse) SetStatusCode(v int32) *ListDataIngestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataIngestionsResponse) SetBody(v *ListDataIngestionsResponseBody) *ListDataIngestionsResponse {
	s.Body = v
	return s
}

func (s *ListDataIngestionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
