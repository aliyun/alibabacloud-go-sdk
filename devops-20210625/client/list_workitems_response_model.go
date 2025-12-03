// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkitemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkitemsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkitemsResponseBody) *ListWorkitemsResponse
	GetBody() *ListWorkitemsResponseBody
}

type ListWorkitemsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkitemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkitemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkitemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkitemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkitemsResponse) GetBody() *ListWorkitemsResponseBody {
	return s.Body
}

func (s *ListWorkitemsResponse) SetHeaders(v map[string]*string) *ListWorkitemsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkitemsResponse) SetStatusCode(v int32) *ListWorkitemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkitemsResponse) SetBody(v *ListWorkitemsResponseBody) *ListWorkitemsResponse {
	s.Body = v
	return s
}

func (s *ListWorkitemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
