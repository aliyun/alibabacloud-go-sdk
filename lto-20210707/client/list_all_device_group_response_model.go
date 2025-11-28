// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListAllDeviceGroupResponseBody) *ListAllDeviceGroupResponse
	GetBody() *ListAllDeviceGroupResponseBody
}

type ListAllDeviceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListAllDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllDeviceGroupResponse) GetBody() *ListAllDeviceGroupResponseBody {
	return s.Body
}

func (s *ListAllDeviceGroupResponse) SetHeaders(v map[string]*string) *ListAllDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListAllDeviceGroupResponse) SetStatusCode(v int32) *ListAllDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllDeviceGroupResponse) SetBody(v *ListAllDeviceGroupResponseBody) *ListAllDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *ListAllDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
