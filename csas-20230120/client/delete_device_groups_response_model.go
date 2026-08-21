// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeviceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeviceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeviceGroupsResponseBody) *DeleteDeviceGroupsResponse
	GetBody() *DeleteDeviceGroupsResponseBody
}

type DeleteDeviceGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeviceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeviceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeviceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeviceGroupsResponse) GetBody() *DeleteDeviceGroupsResponseBody {
	return s.Body
}

func (s *DeleteDeviceGroupsResponse) SetHeaders(v map[string]*string) *DeleteDeviceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceGroupsResponse) SetStatusCode(v int32) *DeleteDeviceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceGroupsResponse) SetBody(v *DeleteDeviceGroupsResponseBody) *DeleteDeviceGroupsResponse {
	s.Body = v
	return s
}

func (s *DeleteDeviceGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
