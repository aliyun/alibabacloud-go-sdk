// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeviceGroupResponseBody) *CreateDeviceGroupResponse
	GetBody() *CreateDeviceGroupResponseBody
}

type CreateDeviceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeviceGroupResponse) GetBody() *CreateDeviceGroupResponseBody {
	return s.Body
}

func (s *CreateDeviceGroupResponse) SetHeaders(v map[string]*string) *CreateDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDeviceGroupResponse) SetStatusCode(v int32) *CreateDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeviceGroupResponse) SetBody(v *CreateDeviceGroupResponseBody) *CreateDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
