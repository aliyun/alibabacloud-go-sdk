// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudConnectNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudConnectNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudConnectNetworkResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudConnectNetworkResponseBody) *CreateCloudConnectNetworkResponse
	GetBody() *CreateCloudConnectNetworkResponseBody
}

type CreateCloudConnectNetworkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudConnectNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudConnectNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudConnectNetworkResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudConnectNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudConnectNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudConnectNetworkResponse) GetBody() *CreateCloudConnectNetworkResponseBody {
	return s.Body
}

func (s *CreateCloudConnectNetworkResponse) SetHeaders(v map[string]*string) *CreateCloudConnectNetworkResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudConnectNetworkResponse) SetStatusCode(v int32) *CreateCloudConnectNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudConnectNetworkResponse) SetBody(v *CreateCloudConnectNetworkResponseBody) *CreateCloudConnectNetworkResponse {
	s.Body = v
	return s
}

func (s *CreateCloudConnectNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
