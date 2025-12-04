// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachElasticNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachElasticNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachElasticNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *AttachElasticNetworkInterfaceResponseBody) *AttachElasticNetworkInterfaceResponse
	GetBody() *AttachElasticNetworkInterfaceResponseBody
}

type AttachElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachElasticNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *AttachElasticNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachElasticNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachElasticNetworkInterfaceResponse) GetBody() *AttachElasticNetworkInterfaceResponseBody {
	return s.Body
}

func (s *AttachElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *AttachElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *AttachElasticNetworkInterfaceResponse) SetStatusCode(v int32) *AttachElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachElasticNetworkInterfaceResponse) SetBody(v *AttachElasticNetworkInterfaceResponseBody) *AttachElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *AttachElasticNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
