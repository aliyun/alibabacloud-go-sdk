// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseProductResponse
	GetStatusCode() *int32
	SetBody(v *CloseProductResponseBody) *CloseProductResponse
	GetBody() *CloseProductResponseBody
}

type CloseProductResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseProductResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseProductResponse) GoString() string {
	return s.String()
}

func (s *CloseProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseProductResponse) GetBody() *CloseProductResponseBody {
	return s.Body
}

func (s *CloseProductResponse) SetHeaders(v map[string]*string) *CloseProductResponse {
	s.Headers = v
	return s
}

func (s *CloseProductResponse) SetStatusCode(v int32) *CloseProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseProductResponse) SetBody(v *CloseProductResponseBody) *CloseProductResponse {
	s.Body = v
	return s
}

func (s *CloseProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
