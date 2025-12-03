// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenProductResponse
	GetStatusCode() *int32
	SetBody(v *OpenProductResponseBody) *OpenProductResponse
	GetBody() *OpenProductResponseBody
}

type OpenProductResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenProductResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenProductResponse) GoString() string {
	return s.String()
}

func (s *OpenProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenProductResponse) GetBody() *OpenProductResponseBody {
	return s.Body
}

func (s *OpenProductResponse) SetHeaders(v map[string]*string) *OpenProductResponse {
	s.Headers = v
	return s
}

func (s *OpenProductResponse) SetStatusCode(v int32) *OpenProductResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenProductResponse) SetBody(v *OpenProductResponseBody) *OpenProductResponse {
	s.Body = v
	return s
}

func (s *OpenProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
