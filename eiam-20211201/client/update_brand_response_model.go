// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBrandResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBrandResponseBody) *UpdateBrandResponse
	GetBody() *UpdateBrandResponseBody
}

type UpdateBrandResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBrandResponse) GoString() string {
	return s.String()
}

func (s *UpdateBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBrandResponse) GetBody() *UpdateBrandResponseBody {
	return s.Body
}

func (s *UpdateBrandResponse) SetHeaders(v map[string]*string) *UpdateBrandResponse {
	s.Headers = v
	return s
}

func (s *UpdateBrandResponse) SetStatusCode(v int32) *UpdateBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBrandResponse) SetBody(v *UpdateBrandResponseBody) *UpdateBrandResponse {
	s.Body = v
	return s
}

func (s *UpdateBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
