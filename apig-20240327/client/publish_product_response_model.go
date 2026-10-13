// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishProductResponse
	GetStatusCode() *int32
	SetBody(v *PublishProductResponseBody) *PublishProductResponse
	GetBody() *PublishProductResponseBody
}

type PublishProductResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishProductResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishProductResponse) GoString() string {
	return s.String()
}

func (s *PublishProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishProductResponse) GetBody() *PublishProductResponseBody {
	return s.Body
}

func (s *PublishProductResponse) SetHeaders(v map[string]*string) *PublishProductResponse {
	s.Headers = v
	return s
}

func (s *PublishProductResponse) SetStatusCode(v int32) *PublishProductResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishProductResponse) SetBody(v *PublishProductResponseBody) *PublishProductResponse {
	s.Body = v
	return s
}

func (s *PublishProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
