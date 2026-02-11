// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllProdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllProdsResponse
	GetStatusCode() *int32
	SetBody(v *ListAllProdsResponseBody) *ListAllProdsResponse
	GetBody() *ListAllProdsResponseBody
}

type ListAllProdsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllProdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllProdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllProdsResponse) GoString() string {
	return s.String()
}

func (s *ListAllProdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllProdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllProdsResponse) GetBody() *ListAllProdsResponseBody {
	return s.Body
}

func (s *ListAllProdsResponse) SetHeaders(v map[string]*string) *ListAllProdsResponse {
	s.Headers = v
	return s
}

func (s *ListAllProdsResponse) SetStatusCode(v int32) *ListAllProdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllProdsResponse) SetBody(v *ListAllProdsResponseBody) *ListAllProdsResponse {
	s.Body = v
	return s
}

func (s *ListAllProdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
