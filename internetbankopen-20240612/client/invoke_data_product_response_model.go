// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeDataProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeDataProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeDataProductResponse
	GetStatusCode() *int32
	SetBody(v *InvokeDataProductResponseBody) *InvokeDataProductResponse
	GetBody() *InvokeDataProductResponseBody
}

type InvokeDataProductResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeDataProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeDataProductResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeDataProductResponse) GoString() string {
	return s.String()
}

func (s *InvokeDataProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeDataProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeDataProductResponse) GetBody() *InvokeDataProductResponseBody {
	return s.Body
}

func (s *InvokeDataProductResponse) SetHeaders(v map[string]*string) *InvokeDataProductResponse {
	s.Headers = v
	return s
}

func (s *InvokeDataProductResponse) SetStatusCode(v int32) *InvokeDataProductResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeDataProductResponse) SetBody(v *InvokeDataProductResponseBody) *InvokeDataProductResponse {
	s.Body = v
	return s
}

func (s *InvokeDataProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
