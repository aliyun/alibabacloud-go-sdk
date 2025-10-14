// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVendorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVendorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVendorsResponse
	GetStatusCode() *int32
	SetBody(v *ListVendorsResponseBody) *ListVendorsResponse
	GetBody() *ListVendorsResponseBody
}

type ListVendorsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVendorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVendorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVendorsResponse) GoString() string {
	return s.String()
}

func (s *ListVendorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVendorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVendorsResponse) GetBody() *ListVendorsResponseBody {
	return s.Body
}

func (s *ListVendorsResponse) SetHeaders(v map[string]*string) *ListVendorsResponse {
	s.Headers = v
	return s
}

func (s *ListVendorsResponse) SetStatusCode(v int32) *ListVendorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVendorsResponse) SetBody(v *ListVendorsResponseBody) *ListVendorsResponse {
	s.Body = v
	return s
}

func (s *ListVendorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
