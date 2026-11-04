// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceDeviceToDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceDeviceToDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceDeviceToDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceDeviceToDeviceGroupResponseBody) *AddFaceDeviceToDeviceGroupResponse
	GetBody() *AddFaceDeviceToDeviceGroupResponseBody
}

type AddFaceDeviceToDeviceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceDeviceToDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceDeviceToDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceDeviceToDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceToDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceDeviceToDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceDeviceToDeviceGroupResponse) GetBody() *AddFaceDeviceToDeviceGroupResponseBody {
	return s.Body
}

func (s *AddFaceDeviceToDeviceGroupResponse) SetHeaders(v map[string]*string) *AddFaceDeviceToDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponse) SetStatusCode(v int32) *AddFaceDeviceToDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponse) SetBody(v *AddFaceDeviceToDeviceGroupResponseBody) *AddFaceDeviceToDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *AddFaceDeviceToDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
