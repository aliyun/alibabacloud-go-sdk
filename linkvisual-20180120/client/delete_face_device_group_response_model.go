// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceDeviceGroupResponseBody) *DeleteFaceDeviceGroupResponse
	GetBody() *DeleteFaceDeviceGroupResponseBody
}

type DeleteFaceDeviceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceDeviceGroupResponse) GetBody() *DeleteFaceDeviceGroupResponseBody {
	return s.Body
}

func (s *DeleteFaceDeviceGroupResponse) SetHeaders(v map[string]*string) *DeleteFaceDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceDeviceGroupResponse) SetStatusCode(v int32) *DeleteFaceDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceDeviceGroupResponse) SetBody(v *DeleteFaceDeviceGroupResponseBody) *DeleteFaceDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
