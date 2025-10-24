// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeQuotaScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeQuotaScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeQuotaScheduleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeQuotaScheduleResponseBody) *UpdateComputeQuotaScheduleResponse
	GetBody() *UpdateComputeQuotaScheduleResponseBody
}

type UpdateComputeQuotaScheduleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeQuotaScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeQuotaScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeQuotaScheduleResponse) GetBody() *UpdateComputeQuotaScheduleResponseBody {
	return s.Body
}

func (s *UpdateComputeQuotaScheduleResponse) SetHeaders(v map[string]*string) *UpdateComputeQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeQuotaScheduleResponse) SetStatusCode(v int32) *UpdateComputeQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeQuotaScheduleResponse) SetBody(v *UpdateComputeQuotaScheduleResponseBody) *UpdateComputeQuotaScheduleResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeQuotaScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
