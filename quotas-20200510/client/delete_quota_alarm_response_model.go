// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQuotaAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQuotaAlarmResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQuotaAlarmResponseBody) *DeleteQuotaAlarmResponse
	GetBody() *DeleteQuotaAlarmResponseBody
}

type DeleteQuotaAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQuotaAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaAlarmResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuotaAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQuotaAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQuotaAlarmResponse) GetBody() *DeleteQuotaAlarmResponseBody {
	return s.Body
}

func (s *DeleteQuotaAlarmResponse) SetHeaders(v map[string]*string) *DeleteQuotaAlarmResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuotaAlarmResponse) SetStatusCode(v int32) *DeleteQuotaAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQuotaAlarmResponse) SetBody(v *DeleteQuotaAlarmResponseBody) *DeleteQuotaAlarmResponse {
	s.Body = v
	return s
}

func (s *DeleteQuotaAlarmResponse) Validate() error {
	return dara.Validate(s)
}
