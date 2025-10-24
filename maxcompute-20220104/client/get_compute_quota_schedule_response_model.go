// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeQuotaScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeQuotaScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeQuotaScheduleResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeQuotaScheduleResponseBody) *GetComputeQuotaScheduleResponse
	GetBody() *GetComputeQuotaScheduleResponseBody
}

type GetComputeQuotaScheduleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeQuotaScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetComputeQuotaScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeQuotaScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeQuotaScheduleResponse) GetBody() *GetComputeQuotaScheduleResponseBody {
	return s.Body
}

func (s *GetComputeQuotaScheduleResponse) SetHeaders(v map[string]*string) *GetComputeQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetComputeQuotaScheduleResponse) SetStatusCode(v int32) *GetComputeQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeQuotaScheduleResponse) SetBody(v *GetComputeQuotaScheduleResponseBody) *GetComputeQuotaScheduleResponse {
	s.Body = v
	return s
}

func (s *GetComputeQuotaScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
