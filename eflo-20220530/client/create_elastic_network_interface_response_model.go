// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateElasticNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateElasticNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateElasticNetworkInterfaceResponseBody) *CreateElasticNetworkInterfaceResponse
	GetBody() *CreateElasticNetworkInterfaceResponseBody
}

type CreateElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateElasticNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateElasticNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateElasticNetworkInterfaceResponse) GetBody() *CreateElasticNetworkInterfaceResponseBody {
	return s.Body
}

func (s *CreateElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *CreateElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticNetworkInterfaceResponse) SetStatusCode(v int32) *CreateElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateElasticNetworkInterfaceResponse) SetBody(v *CreateElasticNetworkInterfaceResponseBody) *CreateElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *CreateElasticNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
