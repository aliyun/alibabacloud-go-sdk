// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubscribeProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubscribeProductResponse
	GetStatusCode() *int32
	SetBody(v *SubscribeProductResponseBody) *SubscribeProductResponse
	GetBody() *SubscribeProductResponseBody
}

type SubscribeProductResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribeProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribeProductResponse) String() string {
	return dara.Prettify(s)
}

func (s SubscribeProductResponse) GoString() string {
	return s.String()
}

func (s *SubscribeProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubscribeProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubscribeProductResponse) GetBody() *SubscribeProductResponseBody {
	return s.Body
}

func (s *SubscribeProductResponse) SetHeaders(v map[string]*string) *SubscribeProductResponse {
	s.Headers = v
	return s
}

func (s *SubscribeProductResponse) SetStatusCode(v int32) *SubscribeProductResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeProductResponse) SetBody(v *SubscribeProductResponseBody) *SubscribeProductResponse {
	s.Body = v
	return s
}

func (s *SubscribeProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
