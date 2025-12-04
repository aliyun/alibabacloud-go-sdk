// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListElasticNetworkInterfacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListElasticNetworkInterfacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListElasticNetworkInterfacesResponse
	GetStatusCode() *int32
	SetBody(v *ListElasticNetworkInterfacesResponseBody) *ListElasticNetworkInterfacesResponse
	GetBody() *ListElasticNetworkInterfacesResponseBody
}

type ListElasticNetworkInterfacesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListElasticNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListElasticNetworkInterfacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListElasticNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListElasticNetworkInterfacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListElasticNetworkInterfacesResponse) GetBody() *ListElasticNetworkInterfacesResponseBody {
	return s.Body
}

func (s *ListElasticNetworkInterfacesResponse) SetHeaders(v map[string]*string) *ListElasticNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *ListElasticNetworkInterfacesResponse) SetStatusCode(v int32) *ListElasticNetworkInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListElasticNetworkInterfacesResponse) SetBody(v *ListElasticNetworkInterfacesResponseBody) *ListElasticNetworkInterfacesResponse {
	s.Body = v
	return s
}

func (s *ListElasticNetworkInterfacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
