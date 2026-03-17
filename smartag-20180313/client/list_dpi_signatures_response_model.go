// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiSignaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDpiSignaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDpiSignaturesResponse
	GetStatusCode() *int32
	SetBody(v *ListDpiSignaturesResponseBody) *ListDpiSignaturesResponse
	GetBody() *ListDpiSignaturesResponseBody
}

type ListDpiSignaturesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDpiSignaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDpiSignaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDpiSignaturesResponse) GoString() string {
	return s.String()
}

func (s *ListDpiSignaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDpiSignaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDpiSignaturesResponse) GetBody() *ListDpiSignaturesResponseBody {
	return s.Body
}

func (s *ListDpiSignaturesResponse) SetHeaders(v map[string]*string) *ListDpiSignaturesResponse {
	s.Headers = v
	return s
}

func (s *ListDpiSignaturesResponse) SetStatusCode(v int32) *ListDpiSignaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDpiSignaturesResponse) SetBody(v *ListDpiSignaturesResponseBody) *ListDpiSignaturesResponse {
	s.Body = v
	return s
}

func (s *ListDpiSignaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
