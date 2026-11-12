// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDisksResponse
	GetStatusCode() *int32
	SetBody(v *ListDisksResponseBody) *ListDisksResponse
	GetBody() *ListDisksResponseBody
}

type ListDisksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDisksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDisksResponse) GoString() string {
	return s.String()
}

func (s *ListDisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDisksResponse) GetBody() *ListDisksResponseBody {
	return s.Body
}

func (s *ListDisksResponse) SetHeaders(v map[string]*string) *ListDisksResponse {
	s.Headers = v
	return s
}

func (s *ListDisksResponse) SetStatusCode(v int32) *ListDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisksResponse) SetBody(v *ListDisksResponseBody) *ListDisksResponse {
	s.Body = v
	return s
}

func (s *ListDisksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
