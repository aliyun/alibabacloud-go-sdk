// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkInterfaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkInterfaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkInterfaceResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkInterfaceResponseBody) *GetNetworkInterfaceResponse
	GetBody() *GetNetworkInterfaceResponseBody
}

type GetNetworkInterfaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkInterfaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkInterfaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkInterfaceResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkInterfaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkInterfaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkInterfaceResponse) GetBody() *GetNetworkInterfaceResponseBody {
	return s.Body
}

func (s *GetNetworkInterfaceResponse) SetHeaders(v map[string]*string) *GetNetworkInterfaceResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkInterfaceResponse) SetStatusCode(v int32) *GetNetworkInterfaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkInterfaceResponse) SetBody(v *GetNetworkInterfaceResponseBody) *GetNetworkInterfaceResponse {
	s.Body = v
	return s
}

func (s *GetNetworkInterfaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
