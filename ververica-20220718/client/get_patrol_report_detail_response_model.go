// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatrolReportDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPatrolReportDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPatrolReportDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetPatrolReportDetailResponseBody) *GetPatrolReportDetailResponse
	GetBody() *GetPatrolReportDetailResponseBody
}

type GetPatrolReportDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPatrolReportDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPatrolReportDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailResponse) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPatrolReportDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPatrolReportDetailResponse) GetBody() *GetPatrolReportDetailResponseBody {
	return s.Body
}

func (s *GetPatrolReportDetailResponse) SetHeaders(v map[string]*string) *GetPatrolReportDetailResponse {
	s.Headers = v
	return s
}

func (s *GetPatrolReportDetailResponse) SetStatusCode(v int32) *GetPatrolReportDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPatrolReportDetailResponse) SetBody(v *GetPatrolReportDetailResponseBody) *GetPatrolReportDetailResponse {
	s.Body = v
	return s
}

func (s *GetPatrolReportDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
