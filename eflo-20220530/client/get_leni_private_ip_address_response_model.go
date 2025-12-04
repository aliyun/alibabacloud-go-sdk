// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLeniPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLeniPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLeniPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *GetLeniPrivateIpAddressResponseBody) *GetLeniPrivateIpAddressResponse
	GetBody() *GetLeniPrivateIpAddressResponseBody
}

type GetLeniPrivateIpAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLeniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLeniPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLeniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *GetLeniPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLeniPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLeniPrivateIpAddressResponse) GetBody() *GetLeniPrivateIpAddressResponseBody {
	return s.Body
}

func (s *GetLeniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *GetLeniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *GetLeniPrivateIpAddressResponse) SetStatusCode(v int32) *GetLeniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLeniPrivateIpAddressResponse) SetBody(v *GetLeniPrivateIpAddressResponseBody) *GetLeniPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *GetLeniPrivateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
