// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserPoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserPoolsResponseBody) *ListUserPoolsResponse
	GetBody() *ListUserPoolsResponseBody
}

type ListUserPoolsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListUserPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserPoolsResponse) GetBody() *ListUserPoolsResponseBody {
	return s.Body
}

func (s *ListUserPoolsResponse) SetHeaders(v map[string]*string) *ListUserPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListUserPoolsResponse) SetStatusCode(v int32) *ListUserPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPoolsResponse) SetBody(v *ListUserPoolsResponseBody) *ListUserPoolsResponse {
	s.Body = v
	return s
}

func (s *ListUserPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
