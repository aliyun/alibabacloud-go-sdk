// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePodsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodePodsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodePodsResponse
	GetStatusCode() *int32
	SetBody(v *ListNodePodsResponseBody) *ListNodePodsResponse
	GetBody() *ListNodePodsResponseBody
}

type ListNodePodsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodePodsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodePodsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodePodsResponse) GoString() string {
	return s.String()
}

func (s *ListNodePodsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodePodsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodePodsResponse) GetBody() *ListNodePodsResponseBody {
	return s.Body
}

func (s *ListNodePodsResponse) SetHeaders(v map[string]*string) *ListNodePodsResponse {
	s.Headers = v
	return s
}

func (s *ListNodePodsResponse) SetStatusCode(v int32) *ListNodePodsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodePodsResponse) SetBody(v *ListNodePodsResponseBody) *ListNodePodsResponse {
	s.Body = v
	return s
}

func (s *ListNodePodsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
