// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCheckProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCheckProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCheckProductResponse
	GetStatusCode() *int32
	SetBody(v *DisableCheckProductResponseBody) *DisableCheckProductResponse
	GetBody() *DisableCheckProductResponseBody
}

type DisableCheckProductResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCheckProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCheckProductResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCheckProductResponse) GoString() string {
	return s.String()
}

func (s *DisableCheckProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCheckProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCheckProductResponse) GetBody() *DisableCheckProductResponseBody {
	return s.Body
}

func (s *DisableCheckProductResponse) SetHeaders(v map[string]*string) *DisableCheckProductResponse {
	s.Headers = v
	return s
}

func (s *DisableCheckProductResponse) SetStatusCode(v int32) *DisableCheckProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCheckProductResponse) SetBody(v *DisableCheckProductResponseBody) *DisableCheckProductResponse {
	s.Body = v
	return s
}

func (s *DisableCheckProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
