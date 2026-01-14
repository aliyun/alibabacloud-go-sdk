// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLogStoreToEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachLogStoreToEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachLogStoreToEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *AttachLogStoreToEndpointGroupResponseBody) *AttachLogStoreToEndpointGroupResponse
	GetBody() *AttachLogStoreToEndpointGroupResponseBody
}

type AttachLogStoreToEndpointGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachLogStoreToEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachLogStoreToEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachLogStoreToEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachLogStoreToEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachLogStoreToEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachLogStoreToEndpointGroupResponse) GetBody() *AttachLogStoreToEndpointGroupResponseBody {
	return s.Body
}

func (s *AttachLogStoreToEndpointGroupResponse) SetHeaders(v map[string]*string) *AttachLogStoreToEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachLogStoreToEndpointGroupResponse) SetStatusCode(v int32) *AttachLogStoreToEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupResponse) SetBody(v *AttachLogStoreToEndpointGroupResponseBody) *AttachLogStoreToEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *AttachLogStoreToEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
