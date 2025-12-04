// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubnetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubnetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubnetsResponse
	GetStatusCode() *int32
	SetBody(v *ListSubnetsResponseBody) *ListSubnetsResponse
	GetBody() *ListSubnetsResponseBody
}

type ListSubnetsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubnetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubnetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubnetsResponse) GoString() string {
	return s.String()
}

func (s *ListSubnetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubnetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubnetsResponse) GetBody() *ListSubnetsResponseBody {
	return s.Body
}

func (s *ListSubnetsResponse) SetHeaders(v map[string]*string) *ListSubnetsResponse {
	s.Headers = v
	return s
}

func (s *ListSubnetsResponse) SetStatusCode(v int32) *ListSubnetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubnetsResponse) SetBody(v *ListSubnetsResponseBody) *ListSubnetsResponse {
	s.Body = v
	return s
}

func (s *ListSubnetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
