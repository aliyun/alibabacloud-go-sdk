// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductResponse
	GetStatusCode() *int32
	SetBody(v *GetProductResponseBody) *GetProductResponse
	GetBody() *GetProductResponseBody
}

type GetProductResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductResponse) GoString() string {
	return s.String()
}

func (s *GetProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductResponse) GetBody() *GetProductResponseBody {
	return s.Body
}

func (s *GetProductResponse) SetHeaders(v map[string]*string) *GetProductResponse {
	s.Headers = v
	return s
}

func (s *GetProductResponse) SetStatusCode(v int32) *GetProductResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductResponse) SetBody(v *GetProductResponseBody) *GetProductResponse {
	s.Body = v
	return s
}

func (s *GetProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
