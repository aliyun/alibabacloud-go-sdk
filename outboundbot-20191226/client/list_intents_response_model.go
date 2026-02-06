// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntentsResponse
	GetStatusCode() *int32
	SetBody(v *ListIntentsResponseBody) *ListIntentsResponse
	GetBody() *ListIntentsResponseBody
}

type ListIntentsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntentsResponse) GoString() string {
	return s.String()
}

func (s *ListIntentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntentsResponse) GetBody() *ListIntentsResponseBody {
	return s.Body
}

func (s *ListIntentsResponse) SetHeaders(v map[string]*string) *ListIntentsResponse {
	s.Headers = v
	return s
}

func (s *ListIntentsResponse) SetStatusCode(v int32) *ListIntentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntentsResponse) SetBody(v *ListIntentsResponseBody) *ListIntentsResponse {
	s.Body = v
	return s
}

func (s *ListIntentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
