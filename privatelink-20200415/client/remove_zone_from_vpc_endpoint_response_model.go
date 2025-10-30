// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveZoneFromVpcEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveZoneFromVpcEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveZoneFromVpcEndpointResponse
	GetStatusCode() *int32
	SetBody(v *RemoveZoneFromVpcEndpointResponseBody) *RemoveZoneFromVpcEndpointResponse
	GetBody() *RemoveZoneFromVpcEndpointResponseBody
}

type RemoveZoneFromVpcEndpointResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveZoneFromVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveZoneFromVpcEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveZoneFromVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *RemoveZoneFromVpcEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveZoneFromVpcEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveZoneFromVpcEndpointResponse) GetBody() *RemoveZoneFromVpcEndpointResponseBody {
	return s.Body
}

func (s *RemoveZoneFromVpcEndpointResponse) SetHeaders(v map[string]*string) *RemoveZoneFromVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *RemoveZoneFromVpcEndpointResponse) SetStatusCode(v int32) *RemoveZoneFromVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointResponse) SetBody(v *RemoveZoneFromVpcEndpointResponseBody) *RemoveZoneFromVpcEndpointResponse {
	s.Body = v
	return s
}

func (s *RemoveZoneFromVpcEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
