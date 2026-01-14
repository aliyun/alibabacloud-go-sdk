// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEndpointGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEndpointGroupsResponse
	GetStatusCode() *int32
	SetBody(v *CreateEndpointGroupsResponseBody) *CreateEndpointGroupsResponse
	GetBody() *CreateEndpointGroupsResponseBody
}

type CreateEndpointGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEndpointGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEndpointGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsResponse) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEndpointGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEndpointGroupsResponse) GetBody() *CreateEndpointGroupsResponseBody {
	return s.Body
}

func (s *CreateEndpointGroupsResponse) SetHeaders(v map[string]*string) *CreateEndpointGroupsResponse {
	s.Headers = v
	return s
}

func (s *CreateEndpointGroupsResponse) SetStatusCode(v int32) *CreateEndpointGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEndpointGroupsResponse) SetBody(v *CreateEndpointGroupsResponseBody) *CreateEndpointGroupsResponse {
	s.Body = v
	return s
}

func (s *CreateEndpointGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
