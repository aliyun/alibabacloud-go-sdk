// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessWarrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessWarrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessWarrantResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessWarrantResponseBody) *CreateAccessWarrantResponse
	GetBody() *CreateAccessWarrantResponseBody
}

type CreateAccessWarrantResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessWarrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessWarrantResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessWarrantResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessWarrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessWarrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessWarrantResponse) GetBody() *CreateAccessWarrantResponseBody {
	return s.Body
}

func (s *CreateAccessWarrantResponse) SetHeaders(v map[string]*string) *CreateAccessWarrantResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessWarrantResponse) SetStatusCode(v int32) *CreateAccessWarrantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessWarrantResponse) SetBody(v *CreateAccessWarrantResponseBody) *CreateAccessWarrantResponse {
	s.Body = v
	return s
}

func (s *CreateAccessWarrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
