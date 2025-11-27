// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDeviceAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDeviceAlarmResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDeviceAlarmResponseBody) *ModifyDeviceAlarmResponse
	GetBody() *ModifyDeviceAlarmResponseBody
}

type ModifyDeviceAlarmResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeviceAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeviceAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceAlarmResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDeviceAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDeviceAlarmResponse) GetBody() *ModifyDeviceAlarmResponseBody {
	return s.Body
}

func (s *ModifyDeviceAlarmResponse) SetHeaders(v map[string]*string) *ModifyDeviceAlarmResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeviceAlarmResponse) SetStatusCode(v int32) *ModifyDeviceAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeviceAlarmResponse) SetBody(v *ModifyDeviceAlarmResponseBody) *ModifyDeviceAlarmResponse {
	s.Body = v
	return s
}

func (s *ModifyDeviceAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
