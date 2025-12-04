// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElasticNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetElasticNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetElasticNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *GetElasticNetworkInterfaceResponseBody) *GetElasticNetworkInterfaceResponse
	GetBody() *GetElasticNetworkInterfaceResponseBody
}

type GetElasticNetworkInterfaceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetElasticNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetElasticNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetElasticNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *GetElasticNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetElasticNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetElasticNetworkInterfaceResponse) GetBody() *GetElasticNetworkInterfaceResponseBody {
	return s.Body
}

func (s *GetElasticNetworkInterfaceResponse) SetHeaders(v map[string]*string) *GetElasticNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *GetElasticNetworkInterfaceResponse) SetStatusCode(v int32) *GetElasticNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetElasticNetworkInterfaceResponse) SetBody(v *GetElasticNetworkInterfaceResponseBody) *GetElasticNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *GetElasticNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
