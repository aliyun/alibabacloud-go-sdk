// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckColumnResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCheckColumnResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCheckColumnResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCheckColumnResultsResponseBody) *ListDataCheckColumnResultsResponse
	GetBody() *ListDataCheckColumnResultsResponseBody
}

type ListDataCheckColumnResultsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCheckColumnResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCheckColumnResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckColumnResultsResponse) GoString() string {
	return s.String()
}

func (s *ListDataCheckColumnResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCheckColumnResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCheckColumnResultsResponse) GetBody() *ListDataCheckColumnResultsResponseBody {
	return s.Body
}

func (s *ListDataCheckColumnResultsResponse) SetHeaders(v map[string]*string) *ListDataCheckColumnResultsResponse {
	s.Headers = v
	return s
}

func (s *ListDataCheckColumnResultsResponse) SetStatusCode(v int32) *ListDataCheckColumnResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCheckColumnResultsResponse) SetBody(v *ListDataCheckColumnResultsResponseBody) *ListDataCheckColumnResultsResponse {
	s.Body = v
	return s
}

func (s *ListDataCheckColumnResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
