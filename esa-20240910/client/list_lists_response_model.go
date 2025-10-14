// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListListsResponse
	GetStatusCode() *int32
	SetBody(v *ListListsResponseBody) *ListListsResponse
	GetBody() *ListListsResponseBody
}

type ListListsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListListsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListListsResponse) GoString() string {
	return s.String()
}

func (s *ListListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListListsResponse) GetBody() *ListListsResponseBody {
	return s.Body
}

func (s *ListListsResponse) SetHeaders(v map[string]*string) *ListListsResponse {
	s.Headers = v
	return s
}

func (s *ListListsResponse) SetStatusCode(v int32) *ListListsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListsResponse) SetBody(v *ListListsResponseBody) *ListListsResponse {
	s.Body = v
	return s
}

func (s *ListListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
