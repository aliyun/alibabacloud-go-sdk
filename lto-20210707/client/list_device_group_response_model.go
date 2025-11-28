// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceGroupResponseBody) *ListDeviceGroupResponse
	GetBody() *ListDeviceGroupResponseBody
}

type ListDeviceGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceGroupResponse) GetBody() *ListDeviceGroupResponseBody {
	return s.Body
}

func (s *ListDeviceGroupResponse) SetHeaders(v map[string]*string) *ListDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceGroupResponse) SetStatusCode(v int32) *ListDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceGroupResponse) SetBody(v *ListDeviceGroupResponseBody) *ListDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *ListDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
