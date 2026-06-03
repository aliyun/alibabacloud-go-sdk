// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigInstanceNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigInstanceNetworkResponse
	GetStatusCode() *int32
	SetBody(v *ConfigInstanceNetworkResponseBody) *ConfigInstanceNetworkResponse
	GetBody() *ConfigInstanceNetworkResponseBody
}

type ConfigInstanceNetworkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceNetworkResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigInstanceNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigInstanceNetworkResponse) GetBody() *ConfigInstanceNetworkResponseBody {
	return s.Body
}

func (s *ConfigInstanceNetworkResponse) SetHeaders(v map[string]*string) *ConfigInstanceNetworkResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceNetworkResponse) SetStatusCode(v int32) *ConfigInstanceNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceNetworkResponse) SetBody(v *ConfigInstanceNetworkResponseBody) *ConfigInstanceNetworkResponse {
	s.Body = v
	return s
}

func (s *ConfigInstanceNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
