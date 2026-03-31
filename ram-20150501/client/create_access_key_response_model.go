// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessKeyResponseBody) *CreateAccessKeyResponse
	GetBody() *CreateAccessKeyResponseBody
}

type CreateAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessKeyResponse) GetBody() *CreateAccessKeyResponseBody {
	return s.Body
}

func (s *CreateAccessKeyResponse) SetHeaders(v map[string]*string) *CreateAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessKeyResponse) SetStatusCode(v int32) *CreateAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessKeyResponse) SetBody(v *CreateAccessKeyResponseBody) *CreateAccessKeyResponse {
	s.Body = v
	return s
}

func (s *CreateAccessKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
