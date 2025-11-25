// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserAlarmConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserAlarmConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserAlarmConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserAlarmConfigResponseBody) *ModifyUserAlarmConfigResponse
	GetBody() *ModifyUserAlarmConfigResponseBody
}

type ModifyUserAlarmConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserAlarmConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserAlarmConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserAlarmConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserAlarmConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserAlarmConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserAlarmConfigResponse) GetBody() *ModifyUserAlarmConfigResponseBody {
	return s.Body
}

func (s *ModifyUserAlarmConfigResponse) SetHeaders(v map[string]*string) *ModifyUserAlarmConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserAlarmConfigResponse) SetStatusCode(v int32) *ModifyUserAlarmConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserAlarmConfigResponse) SetBody(v *ModifyUserAlarmConfigResponseBody) *ModifyUserAlarmConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyUserAlarmConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
