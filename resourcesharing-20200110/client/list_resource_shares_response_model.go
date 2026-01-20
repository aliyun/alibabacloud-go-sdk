// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceSharesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceSharesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceSharesResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceSharesResponseBody) *ListResourceSharesResponse
	GetBody() *ListResourceSharesResponseBody
}

type ListResourceSharesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceSharesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceSharesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceSharesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceSharesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceSharesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceSharesResponse) GetBody() *ListResourceSharesResponseBody {
	return s.Body
}

func (s *ListResourceSharesResponse) SetHeaders(v map[string]*string) *ListResourceSharesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceSharesResponse) SetStatusCode(v int32) *ListResourceSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceSharesResponse) SetBody(v *ListResourceSharesResponseBody) *ListResourceSharesResponse {
	s.Body = v
	return s
}

func (s *ListResourceSharesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
