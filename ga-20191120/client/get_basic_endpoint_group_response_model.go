// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicEndpointGroupResponseBody) *GetBasicEndpointGroupResponse
	GetBody() *GetBasicEndpointGroupResponseBody
}

type GetBasicEndpointGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicEndpointGroupResponse) GetBody() *GetBasicEndpointGroupResponseBody {
	return s.Body
}

func (s *GetBasicEndpointGroupResponse) SetHeaders(v map[string]*string) *GetBasicEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *GetBasicEndpointGroupResponse) SetStatusCode(v int32) *GetBasicEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicEndpointGroupResponse) SetBody(v *GetBasicEndpointGroupResponseBody) *GetBasicEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *GetBasicEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
