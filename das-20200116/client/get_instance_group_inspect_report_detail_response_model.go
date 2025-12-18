// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGroupInspectReportDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceGroupInspectReportDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceGroupInspectReportDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceGroupInspectReportDetailResponseBody) *GetInstanceGroupInspectReportDetailResponse
	GetBody() *GetInstanceGroupInspectReportDetailResponseBody
}

type GetInstanceGroupInspectReportDetailResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceGroupInspectReportDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceGroupInspectReportDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportDetailResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceGroupInspectReportDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceGroupInspectReportDetailResponse) GetBody() *GetInstanceGroupInspectReportDetailResponseBody {
	return s.Body
}

func (s *GetInstanceGroupInspectReportDetailResponse) SetHeaders(v map[string]*string) *GetInstanceGroupInspectReportDetailResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponse) SetStatusCode(v int32) *GetInstanceGroupInspectReportDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponse) SetBody(v *GetInstanceGroupInspectReportDetailResponseBody) *GetInstanceGroupInspectReportDetailResponse {
	s.Body = v
	return s
}

func (s *GetInstanceGroupInspectReportDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
