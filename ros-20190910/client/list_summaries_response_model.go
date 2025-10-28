// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSummariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSummariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSummariesResponse
	GetStatusCode() *int32
	SetBody(v *ListSummariesResponseBody) *ListSummariesResponse
	GetBody() *ListSummariesResponseBody
}

type ListSummariesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSummariesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSummariesResponse) GoString() string {
	return s.String()
}

func (s *ListSummariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSummariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSummariesResponse) GetBody() *ListSummariesResponseBody {
	return s.Body
}

func (s *ListSummariesResponse) SetHeaders(v map[string]*string) *ListSummariesResponse {
	s.Headers = v
	return s
}

func (s *ListSummariesResponse) SetStatusCode(v int32) *ListSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSummariesResponse) SetBody(v *ListSummariesResponseBody) *ListSummariesResponse {
	s.Body = v
	return s
}

func (s *ListSummariesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
