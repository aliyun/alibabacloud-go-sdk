// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDockerhubImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDockerhubImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDockerhubImageResponse
	GetStatusCode() *int32
	SetBody(v *ListDockerhubImageResponseBody) *ListDockerhubImageResponse
	GetBody() *ListDockerhubImageResponseBody
}

type ListDockerhubImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDockerhubImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDockerhubImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDockerhubImageResponse) GoString() string {
	return s.String()
}

func (s *ListDockerhubImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDockerhubImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDockerhubImageResponse) GetBody() *ListDockerhubImageResponseBody {
	return s.Body
}

func (s *ListDockerhubImageResponse) SetHeaders(v map[string]*string) *ListDockerhubImageResponse {
	s.Headers = v
	return s
}

func (s *ListDockerhubImageResponse) SetStatusCode(v int32) *ListDockerhubImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDockerhubImageResponse) SetBody(v *ListDockerhubImageResponseBody) *ListDockerhubImageResponse {
	s.Body = v
	return s
}

func (s *ListDockerhubImageResponse) Validate() error {
	return dara.Validate(s)
}
