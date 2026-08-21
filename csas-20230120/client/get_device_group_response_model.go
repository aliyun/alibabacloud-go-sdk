// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceGroupResponseBody) *GetDeviceGroupResponse
	GetBody() *GetDeviceGroupResponseBody
}

type GetDeviceGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceGroupResponse) GetBody() *GetDeviceGroupResponseBody {
	return s.Body
}

func (s *GetDeviceGroupResponse) SetHeaders(v map[string]*string) *GetDeviceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceGroupResponse) SetStatusCode(v int32) *GetDeviceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceGroupResponse) SetBody(v *GetDeviceGroupResponseBody) *GetDeviceGroupResponse {
	s.Body = v
	return s
}

func (s *GetDeviceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
