// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkServiceResponseBody) *CreateNetworkServiceResponse
	GetBody() *CreateNetworkServiceResponseBody
}

type CreateNetworkServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkServiceResponse) GetBody() *CreateNetworkServiceResponseBody {
	return s.Body
}

func (s *CreateNetworkServiceResponse) SetHeaders(v map[string]*string) *CreateNetworkServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkServiceResponse) SetStatusCode(v int32) *CreateNetworkServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkServiceResponse) SetBody(v *CreateNetworkServiceResponseBody) *CreateNetworkServiceResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
