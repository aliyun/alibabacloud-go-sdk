// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlarmResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlarmResponseBody) *DeleteAlarmResponse
	GetBody() *DeleteAlarmResponseBody
}

type DeleteAlarmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlarmResponse) GetBody() *DeleteAlarmResponseBody {
	return s.Body
}

func (s *DeleteAlarmResponse) SetHeaders(v map[string]*string) *DeleteAlarmResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlarmResponse) SetStatusCode(v int32) *DeleteAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlarmResponse) SetBody(v *DeleteAlarmResponseBody) *DeleteAlarmResponse {
	s.Body = v
	return s
}

func (s *DeleteAlarmResponse) Validate() error {
	return dara.Validate(s)
}
