// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLibrarysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLibrarysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLibrarysResponse
	GetStatusCode() *int32
	SetBody(v *ListLibrarysResponseBody) *ListLibrarysResponse
	GetBody() *ListLibrarysResponseBody
}

type ListLibrarysResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLibrarysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLibrarysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLibrarysResponse) GoString() string {
	return s.String()
}

func (s *ListLibrarysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLibrarysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLibrarysResponse) GetBody() *ListLibrarysResponseBody {
	return s.Body
}

func (s *ListLibrarysResponse) SetHeaders(v map[string]*string) *ListLibrarysResponse {
	s.Headers = v
	return s
}

func (s *ListLibrarysResponse) SetStatusCode(v int32) *ListLibrarysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLibrarysResponse) SetBody(v *ListLibrarysResponseBody) *ListLibrarysResponse {
	s.Body = v
	return s
}

func (s *ListLibrarysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
