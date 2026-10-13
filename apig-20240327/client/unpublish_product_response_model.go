// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnpublishProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnpublishProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnpublishProductResponse
	GetStatusCode() *int32
	SetBody(v *UnpublishProductResponseBody) *UnpublishProductResponse
	GetBody() *UnpublishProductResponseBody
}

type UnpublishProductResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnpublishProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnpublishProductResponse) String() string {
	return dara.Prettify(s)
}

func (s UnpublishProductResponse) GoString() string {
	return s.String()
}

func (s *UnpublishProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnpublishProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnpublishProductResponse) GetBody() *UnpublishProductResponseBody {
	return s.Body
}

func (s *UnpublishProductResponse) SetHeaders(v map[string]*string) *UnpublishProductResponse {
	s.Headers = v
	return s
}

func (s *UnpublishProductResponse) SetStatusCode(v int32) *UnpublishProductResponse {
	s.StatusCode = &v
	return s
}

func (s *UnpublishProductResponse) SetBody(v *UnpublishProductResponseBody) *UnpublishProductResponse {
	s.Body = v
	return s
}

func (s *UnpublishProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
