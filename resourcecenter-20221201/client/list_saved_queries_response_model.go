// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSavedQueriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSavedQueriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSavedQueriesResponse
	GetStatusCode() *int32
	SetBody(v *ListSavedQueriesResponseBody) *ListSavedQueriesResponse
	GetBody() *ListSavedQueriesResponseBody
}

type ListSavedQueriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSavedQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSavedQueriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSavedQueriesResponse) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSavedQueriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSavedQueriesResponse) GetBody() *ListSavedQueriesResponseBody {
	return s.Body
}

func (s *ListSavedQueriesResponse) SetHeaders(v map[string]*string) *ListSavedQueriesResponse {
	s.Headers = v
	return s
}

func (s *ListSavedQueriesResponse) SetStatusCode(v int32) *ListSavedQueriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSavedQueriesResponse) SetBody(v *ListSavedQueriesResponseBody) *ListSavedQueriesResponse {
	s.Body = v
	return s
}

func (s *ListSavedQueriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
