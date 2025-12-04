// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLniPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLniPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLniPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *GetLniPrivateIpAddressResponseBody) *GetLniPrivateIpAddressResponse
	GetBody() *GetLniPrivateIpAddressResponseBody
}

type GetLniPrivateIpAddressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLniPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *GetLniPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLniPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLniPrivateIpAddressResponse) GetBody() *GetLniPrivateIpAddressResponseBody {
	return s.Body
}

func (s *GetLniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *GetLniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *GetLniPrivateIpAddressResponse) SetStatusCode(v int32) *GetLniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLniPrivateIpAddressResponse) SetBody(v *GetLniPrivateIpAddressResponseBody) *GetLniPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *GetLniPrivateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
