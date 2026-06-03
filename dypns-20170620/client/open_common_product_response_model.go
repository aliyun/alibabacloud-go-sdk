// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCommonProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenCommonProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenCommonProductResponse
	GetStatusCode() *int32
	SetBody(v *OpenCommonProductResponseBody) *OpenCommonProductResponse
	GetBody() *OpenCommonProductResponseBody
}

type OpenCommonProductResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenCommonProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenCommonProductResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenCommonProductResponse) GoString() string {
	return s.String()
}

func (s *OpenCommonProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenCommonProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenCommonProductResponse) GetBody() *OpenCommonProductResponseBody {
	return s.Body
}

func (s *OpenCommonProductResponse) SetHeaders(v map[string]*string) *OpenCommonProductResponse {
	s.Headers = v
	return s
}

func (s *OpenCommonProductResponse) SetStatusCode(v int32) *OpenCommonProductResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCommonProductResponse) SetBody(v *OpenCommonProductResponseBody) *OpenCommonProductResponse {
	s.Body = v
	return s
}

func (s *OpenCommonProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
