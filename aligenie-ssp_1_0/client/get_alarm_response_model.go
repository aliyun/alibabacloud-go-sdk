// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlarmResponse
	GetStatusCode() *int32
	SetBody(v *GetAlarmResponseBody) *GetAlarmResponse
	GetBody() *GetAlarmResponseBody
}

type GetAlarmResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmResponse) GoString() string {
	return s.String()
}

func (s *GetAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlarmResponse) GetBody() *GetAlarmResponseBody {
	return s.Body
}

func (s *GetAlarmResponse) SetHeaders(v map[string]*string) *GetAlarmResponse {
	s.Headers = v
	return s
}

func (s *GetAlarmResponse) SetStatusCode(v int32) *GetAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlarmResponse) SetBody(v *GetAlarmResponseBody) *GetAlarmResponse {
	s.Body = v
	return s
}

func (s *GetAlarmResponse) Validate() error {
	return dara.Validate(s)
}
