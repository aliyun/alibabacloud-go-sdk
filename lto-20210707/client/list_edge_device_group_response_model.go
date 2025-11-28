// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeDeviceGroupResponseBody) *ListEdgeDeviceGroupResponse
	GetBody() *ListEdgeDeviceGroupResponseBody
}

type ListEdgeDeviceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeDeviceGroupResponse) GetBody() *ListEdgeDeviceGroupResponseBody {
	return s.Body
}

func (s *ListEdgeDeviceGroupResponse) SetHeaders(v map[string]*string) *ListEdgeDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeDeviceGroupResponse) SetStatusCode(v int32) *ListEdgeDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeDeviceGroupResponse) SetBody(v *ListEdgeDeviceGroupResponseBody) *ListEdgeDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *ListEdgeDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
