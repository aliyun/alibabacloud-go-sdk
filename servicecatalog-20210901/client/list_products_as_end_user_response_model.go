// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsAsEndUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductsAsEndUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductsAsEndUserResponse
	GetStatusCode() *int32
	SetBody(v *ListProductsAsEndUserResponseBody) *ListProductsAsEndUserResponse
	GetBody() *ListProductsAsEndUserResponseBody
}

type ListProductsAsEndUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductsAsEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductsAsEndUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsEndUserResponse) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductsAsEndUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductsAsEndUserResponse) GetBody() *ListProductsAsEndUserResponseBody {
	return s.Body
}

func (s *ListProductsAsEndUserResponse) SetHeaders(v map[string]*string) *ListProductsAsEndUserResponse {
	s.Headers = v
	return s
}

func (s *ListProductsAsEndUserResponse) SetStatusCode(v int32) *ListProductsAsEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductsAsEndUserResponse) SetBody(v *ListProductsAsEndUserResponseBody) *ListProductsAsEndUserResponse {
	s.Body = v
	return s
}

func (s *ListProductsAsEndUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
