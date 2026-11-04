// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceAllDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceAllDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceAllDeviceGroupResponseBody) *QueryFaceAllDeviceGroupResponse
	GetBody() *QueryFaceAllDeviceGroupResponseBody
}

type QueryFaceAllDeviceGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceAllDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceAllDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceAllDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceAllDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceAllDeviceGroupResponse) GetBody() *QueryFaceAllDeviceGroupResponseBody {
	return s.Body
}

func (s *QueryFaceAllDeviceGroupResponse) SetHeaders(v map[string]*string) *QueryFaceAllDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceAllDeviceGroupResponse) SetStatusCode(v int32) *QueryFaceAllDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceAllDeviceGroupResponse) SetBody(v *QueryFaceAllDeviceGroupResponseBody) *QueryFaceAllDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *QueryFaceAllDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
