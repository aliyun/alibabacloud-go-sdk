// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterSubnetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigClusterSubnetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigClusterSubnetResponse
	GetStatusCode() *int32
	SetBody(v *ConfigClusterSubnetResponseBody) *ConfigClusterSubnetResponse
	GetBody() *ConfigClusterSubnetResponseBody
}

type ConfigClusterSubnetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterSubnetResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterSubnetResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigClusterSubnetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigClusterSubnetResponse) GetBody() *ConfigClusterSubnetResponseBody {
	return s.Body
}

func (s *ConfigClusterSubnetResponse) SetHeaders(v map[string]*string) *ConfigClusterSubnetResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterSubnetResponse) SetStatusCode(v int32) *ConfigClusterSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterSubnetResponse) SetBody(v *ConfigClusterSubnetResponseBody) *ConfigClusterSubnetResponse {
	s.Body = v
	return s
}

func (s *ConfigClusterSubnetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
