// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaMarksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaMarksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaMarksResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaMarksResponseBody) *ListMediaMarksResponse
	GetBody() *ListMediaMarksResponseBody
}

type ListMediaMarksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaMarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaMarksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaMarksResponse) GoString() string {
	return s.String()
}

func (s *ListMediaMarksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaMarksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaMarksResponse) GetBody() *ListMediaMarksResponseBody {
	return s.Body
}

func (s *ListMediaMarksResponse) SetHeaders(v map[string]*string) *ListMediaMarksResponse {
	s.Headers = v
	return s
}

func (s *ListMediaMarksResponse) SetStatusCode(v int32) *ListMediaMarksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaMarksResponse) SetBody(v *ListMediaMarksResponseBody) *ListMediaMarksResponse {
	s.Body = v
	return s
}

func (s *ListMediaMarksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
