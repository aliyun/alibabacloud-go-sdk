// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLogStoreFromEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachLogStoreFromEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachLogStoreFromEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *DetachLogStoreFromEndpointGroupResponseBody) *DetachLogStoreFromEndpointGroupResponse
	GetBody() *DetachLogStoreFromEndpointGroupResponseBody
}

type DetachLogStoreFromEndpointGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachLogStoreFromEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachLogStoreFromEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachLogStoreFromEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachLogStoreFromEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachLogStoreFromEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachLogStoreFromEndpointGroupResponse) GetBody() *DetachLogStoreFromEndpointGroupResponseBody {
	return s.Body
}

func (s *DetachLogStoreFromEndpointGroupResponse) SetHeaders(v map[string]*string) *DetachLogStoreFromEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachLogStoreFromEndpointGroupResponse) SetStatusCode(v int32) *DetachLogStoreFromEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupResponse) SetBody(v *DetachLogStoreFromEndpointGroupResponseBody) *DetachLogStoreFromEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *DetachLogStoreFromEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
