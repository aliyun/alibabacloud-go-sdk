// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteElasticNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteElasticNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteElasticNetworkInterfaceResponseBody) *DeleteElasticNetworkInterfaceResponse
	GetBody() *DeleteElasticNetworkInterfaceResponseBody
}

type DeleteElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteElasticNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteElasticNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteElasticNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteElasticNetworkInterfaceResponse) GetBody() *DeleteElasticNetworkInterfaceResponseBody {
	return s.Body
}

func (s *DeleteElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *DeleteElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponse) SetStatusCode(v int32) *DeleteElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponse) SetBody(v *DeleteElasticNetworkInterfaceResponseBody) *DeleteElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *DeleteElasticNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
