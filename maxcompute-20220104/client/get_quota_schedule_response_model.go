// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQuotaScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQuotaScheduleResponse
	GetStatusCode() *int32
	SetBody(v *GetQuotaScheduleResponseBody) *GetQuotaScheduleResponse
	GetBody() *GetQuotaScheduleResponseBody
}

type GetQuotaScheduleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQuotaScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQuotaScheduleResponse) GetBody() *GetQuotaScheduleResponseBody {
	return s.Body
}

func (s *GetQuotaScheduleResponse) SetHeaders(v map[string]*string) *GetQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaScheduleResponse) SetStatusCode(v int32) *GetQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaScheduleResponse) SetBody(v *GetQuotaScheduleResponseBody) *GetQuotaScheduleResponse {
	s.Body = v
	return s
}

func (s *GetQuotaScheduleResponse) Validate() error {
	return dara.Validate(s)
}
