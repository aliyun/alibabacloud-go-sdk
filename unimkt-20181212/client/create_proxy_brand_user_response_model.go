// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyBrandUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProxyBrandUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProxyBrandUserResponse
	GetStatusCode() *int32
	SetBody(v *CreateProxyBrandUserResponseBody) *CreateProxyBrandUserResponse
	GetBody() *CreateProxyBrandUserResponseBody
}

type CreateProxyBrandUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProxyBrandUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProxyBrandUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyBrandUserResponse) GoString() string {
	return s.String()
}

func (s *CreateProxyBrandUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProxyBrandUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProxyBrandUserResponse) GetBody() *CreateProxyBrandUserResponseBody {
	return s.Body
}

func (s *CreateProxyBrandUserResponse) SetHeaders(v map[string]*string) *CreateProxyBrandUserResponse {
	s.Headers = v
	return s
}

func (s *CreateProxyBrandUserResponse) SetStatusCode(v int32) *CreateProxyBrandUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProxyBrandUserResponse) SetBody(v *CreateProxyBrandUserResponseBody) *CreateProxyBrandUserResponse {
	s.Body = v
	return s
}

func (s *CreateProxyBrandUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
