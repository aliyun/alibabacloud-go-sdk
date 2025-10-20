// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBrandResponse
	GetStatusCode() *int32
	SetBody(v *CreateBrandResponseBody) *CreateBrandResponse
	GetBody() *CreateBrandResponseBody
}

type CreateBrandResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBrandResponse) GoString() string {
	return s.String()
}

func (s *CreateBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBrandResponse) GetBody() *CreateBrandResponseBody {
	return s.Body
}

func (s *CreateBrandResponse) SetHeaders(v map[string]*string) *CreateBrandResponse {
	s.Headers = v
	return s
}

func (s *CreateBrandResponse) SetStatusCode(v int32) *CreateBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBrandResponse) SetBody(v *CreateBrandResponseBody) *CreateBrandResponse {
	s.Body = v
	return s
}

func (s *CreateBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
