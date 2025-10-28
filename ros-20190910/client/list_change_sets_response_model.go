// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChangeSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChangeSetsResponse
	GetStatusCode() *int32
	SetBody(v *ListChangeSetsResponseBody) *ListChangeSetsResponse
	GetBody() *ListChangeSetsResponseBody
}

type ListChangeSetsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChangeSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChangeSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChangeSetsResponse) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChangeSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChangeSetsResponse) GetBody() *ListChangeSetsResponseBody {
	return s.Body
}

func (s *ListChangeSetsResponse) SetHeaders(v map[string]*string) *ListChangeSetsResponse {
	s.Headers = v
	return s
}

func (s *ListChangeSetsResponse) SetStatusCode(v int32) *ListChangeSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChangeSetsResponse) SetBody(v *ListChangeSetsResponseBody) *ListChangeSetsResponse {
	s.Body = v
	return s
}

func (s *ListChangeSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
