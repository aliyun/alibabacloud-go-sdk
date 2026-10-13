// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListMemoriesResponseBody) *ListMemoriesResponse
	GetBody() *ListMemoriesResponseBody
}

type ListMemoriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesResponse) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemoriesResponse) GetBody() *ListMemoriesResponseBody {
	return s.Body
}

func (s *ListMemoriesResponse) SetHeaders(v map[string]*string) *ListMemoriesResponse {
	s.Headers = v
	return s
}

func (s *ListMemoriesResponse) SetStatusCode(v int32) *ListMemoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoriesResponse) SetBody(v *ListMemoriesResponseBody) *ListMemoriesResponse {
	s.Body = v
	return s
}

func (s *ListMemoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
