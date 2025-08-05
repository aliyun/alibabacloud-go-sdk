// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQuotaScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQuotaScheduleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQuotaScheduleResponseBody) *UpdateQuotaScheduleResponse
	GetBody() *UpdateQuotaScheduleResponseBody
}

type UpdateQuotaScheduleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQuotaScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQuotaScheduleResponse) GetBody() *UpdateQuotaScheduleResponseBody {
	return s.Body
}

func (s *UpdateQuotaScheduleResponse) SetHeaders(v map[string]*string) *UpdateQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaScheduleResponse) SetStatusCode(v int32) *UpdateQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaScheduleResponse) SetBody(v *UpdateQuotaScheduleResponseBody) *UpdateQuotaScheduleResponse {
	s.Body = v
	return s
}

func (s *UpdateQuotaScheduleResponse) Validate() error {
	return dara.Validate(s)
}
