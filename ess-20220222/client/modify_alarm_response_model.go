// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAlarmResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAlarmResponseBody) *ModifyAlarmResponse
	GetBody() *ModifyAlarmResponseBody
}

type ModifyAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAlarmResponse) GoString() string {
	return s.String()
}

func (s *ModifyAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAlarmResponse) GetBody() *ModifyAlarmResponseBody {
	return s.Body
}

func (s *ModifyAlarmResponse) SetHeaders(v map[string]*string) *ModifyAlarmResponse {
	s.Headers = v
	return s
}

func (s *ModifyAlarmResponse) SetStatusCode(v int32) *ModifyAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAlarmResponse) SetBody(v *ModifyAlarmResponseBody) *ModifyAlarmResponse {
	s.Body = v
	return s
}

func (s *ModifyAlarmResponse) Validate() error {
	return dara.Validate(s)
}
