// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupIdResponseBody) *ListGroupIdResponse
	GetBody() *ListGroupIdResponseBody
}

type ListGroupIdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupIdResponse) GetBody() *ListGroupIdResponseBody {
	return s.Body
}

func (s *ListGroupIdResponse) SetHeaders(v map[string]*string) *ListGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListGroupIdResponse) SetStatusCode(v int32) *ListGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupIdResponse) SetBody(v *ListGroupIdResponseBody) *ListGroupIdResponse {
	s.Body = v
	return s
}

func (s *ListGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
