// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQuotaAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQuotaAlarmResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQuotaAlarmResponseBody) *UpdateQuotaAlarmResponse
	GetBody() *UpdateQuotaAlarmResponseBody
}

type UpdateQuotaAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaAlarmResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQuotaAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQuotaAlarmResponse) GetBody() *UpdateQuotaAlarmResponseBody {
	return s.Body
}

func (s *UpdateQuotaAlarmResponse) SetHeaders(v map[string]*string) *UpdateQuotaAlarmResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaAlarmResponse) SetStatusCode(v int32) *UpdateQuotaAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaAlarmResponse) SetBody(v *UpdateQuotaAlarmResponseBody) *UpdateQuotaAlarmResponse {
	s.Body = v
	return s
}

func (s *UpdateQuotaAlarmResponse) Validate() error {
	return dara.Validate(s)
}
