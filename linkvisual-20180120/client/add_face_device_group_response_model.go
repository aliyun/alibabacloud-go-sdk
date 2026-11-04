// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceDeviceGroupResponseBody) *AddFaceDeviceGroupResponse
	GetBody() *AddFaceDeviceGroupResponseBody
}

type AddFaceDeviceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceDeviceGroupResponse) GetBody() *AddFaceDeviceGroupResponseBody {
	return s.Body
}

func (s *AddFaceDeviceGroupResponse) SetHeaders(v map[string]*string) *AddFaceDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *AddFaceDeviceGroupResponse) SetStatusCode(v int32) *AddFaceDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceDeviceGroupResponse) SetBody(v *AddFaceDeviceGroupResponseBody) *AddFaceDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *AddFaceDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
