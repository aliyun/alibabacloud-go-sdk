// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAlarmResponse
	GetStatusCode() *int32
	SetBody(v *DisableAlarmResponseBody) *DisableAlarmResponse
	GetBody() *DisableAlarmResponseBody
}

type DisableAlarmResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAlarmResponse) GoString() string {
	return s.String()
}

func (s *DisableAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAlarmResponse) GetBody() *DisableAlarmResponseBody {
	return s.Body
}

func (s *DisableAlarmResponse) SetHeaders(v map[string]*string) *DisableAlarmResponse {
	s.Headers = v
	return s
}

func (s *DisableAlarmResponse) SetStatusCode(v int32) *DisableAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAlarmResponse) SetBody(v *DisableAlarmResponseBody) *DisableAlarmResponse {
	s.Body = v
	return s
}

func (s *DisableAlarmResponse) Validate() error {
	return dara.Validate(s)
}
