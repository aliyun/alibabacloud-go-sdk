// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbstractEcommerceVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbstractEcommerceVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbstractEcommerceVideoResponse
	GetStatusCode() *int32
	SetBody(v *AbstractEcommerceVideoResponseBody) *AbstractEcommerceVideoResponse
	GetBody() *AbstractEcommerceVideoResponseBody
}

type AbstractEcommerceVideoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbstractEcommerceVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbstractEcommerceVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s AbstractEcommerceVideoResponse) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbstractEcommerceVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbstractEcommerceVideoResponse) GetBody() *AbstractEcommerceVideoResponseBody {
	return s.Body
}

func (s *AbstractEcommerceVideoResponse) SetHeaders(v map[string]*string) *AbstractEcommerceVideoResponse {
	s.Headers = v
	return s
}

func (s *AbstractEcommerceVideoResponse) SetStatusCode(v int32) *AbstractEcommerceVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *AbstractEcommerceVideoResponse) SetBody(v *AbstractEcommerceVideoResponseBody) *AbstractEcommerceVideoResponse {
	s.Body = v
	return s
}

func (s *AbstractEcommerceVideoResponse) Validate() error {
	return dara.Validate(s)
}
