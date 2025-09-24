// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserAlarmThresholdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetResellerUserAlarmThresholdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetResellerUserAlarmThresholdResponse
	GetStatusCode() *int32
	SetBody(v *SetResellerUserAlarmThresholdResponseBody) *SetResellerUserAlarmThresholdResponse
	GetBody() *SetResellerUserAlarmThresholdResponseBody
}

type SetResellerUserAlarmThresholdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetResellerUserAlarmThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetResellerUserAlarmThresholdResponse) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserAlarmThresholdResponse) GoString() string {
	return s.String()
}

func (s *SetResellerUserAlarmThresholdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetResellerUserAlarmThresholdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetResellerUserAlarmThresholdResponse) GetBody() *SetResellerUserAlarmThresholdResponseBody {
	return s.Body
}

func (s *SetResellerUserAlarmThresholdResponse) SetHeaders(v map[string]*string) *SetResellerUserAlarmThresholdResponse {
	s.Headers = v
	return s
}

func (s *SetResellerUserAlarmThresholdResponse) SetStatusCode(v int32) *SetResellerUserAlarmThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponse) SetBody(v *SetResellerUserAlarmThresholdResponseBody) *SetResellerUserAlarmThresholdResponse {
	s.Body = v
	return s
}

func (s *SetResellerUserAlarmThresholdResponse) Validate() error {
	return dara.Validate(s)
}
