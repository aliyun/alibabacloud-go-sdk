// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateElasticNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateElasticNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateElasticNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateElasticNetworkInterfaceResponseBody) *UpdateElasticNetworkInterfaceResponse
	GetBody() *UpdateElasticNetworkInterfaceResponseBody
}

type UpdateElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateElasticNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateElasticNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateElasticNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateElasticNetworkInterfaceResponse) GetBody() *UpdateElasticNetworkInterfaceResponseBody {
	return s.Body
}

func (s *UpdateElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *UpdateElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponse) SetStatusCode(v int32) *UpdateElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponse) SetBody(v *UpdateElasticNetworkInterfaceResponseBody) *UpdateElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *UpdateElasticNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
