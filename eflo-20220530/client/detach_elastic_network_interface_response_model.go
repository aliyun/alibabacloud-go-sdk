// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachElasticNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachElasticNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachElasticNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *DetachElasticNetworkInterfaceResponseBody) *DetachElasticNetworkInterfaceResponse
	GetBody() *DetachElasticNetworkInterfaceResponseBody
}

type DetachElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachElasticNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *DetachElasticNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachElasticNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachElasticNetworkInterfaceResponse) GetBody() *DetachElasticNetworkInterfaceResponseBody {
	return s.Body
}

func (s *DetachElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *DetachElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *DetachElasticNetworkInterfaceResponse) SetStatusCode(v int32) *DetachElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachElasticNetworkInterfaceResponse) SetBody(v *DetachElasticNetworkInterfaceResponseBody) *DetachElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *DetachElasticNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
