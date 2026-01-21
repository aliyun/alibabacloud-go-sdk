// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductBindBsnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProductBindBsnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProductBindBsnResponse
	GetStatusCode() *int32
	SetBody(v *ProductBindBsnResponseBody) *ProductBindBsnResponse
	GetBody() *ProductBindBsnResponseBody
}

type ProductBindBsnResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductBindBsnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProductBindBsnResponse) String() string {
	return dara.Prettify(s)
}

func (s ProductBindBsnResponse) GoString() string {
	return s.String()
}

func (s *ProductBindBsnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProductBindBsnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProductBindBsnResponse) GetBody() *ProductBindBsnResponseBody {
	return s.Body
}

func (s *ProductBindBsnResponse) SetHeaders(v map[string]*string) *ProductBindBsnResponse {
	s.Headers = v
	return s
}

func (s *ProductBindBsnResponse) SetStatusCode(v int32) *ProductBindBsnResponse {
	s.StatusCode = &v
	return s
}

func (s *ProductBindBsnResponse) SetBody(v *ProductBindBsnResponseBody) *ProductBindBsnResponse {
	s.Body = v
	return s
}

func (s *ProductBindBsnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
