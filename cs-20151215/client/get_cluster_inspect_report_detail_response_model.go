// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterInspectReportDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterInspectReportDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterInspectReportDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterInspectReportDetailResponseBody) *GetClusterInspectReportDetailResponse
	GetBody() *GetClusterInspectReportDetailResponseBody
}

type GetClusterInspectReportDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterInspectReportDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterInspectReportDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterInspectReportDetailResponse) GoString() string {
	return s.String()
}

func (s *GetClusterInspectReportDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterInspectReportDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterInspectReportDetailResponse) GetBody() *GetClusterInspectReportDetailResponseBody {
	return s.Body
}

func (s *GetClusterInspectReportDetailResponse) SetHeaders(v map[string]*string) *GetClusterInspectReportDetailResponse {
	s.Headers = v
	return s
}

func (s *GetClusterInspectReportDetailResponse) SetStatusCode(v int32) *GetClusterInspectReportDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterInspectReportDetailResponse) SetBody(v *GetClusterInspectReportDetailResponseBody) *GetClusterInspectReportDetailResponse {
	s.Body = v
	return s
}

func (s *GetClusterInspectReportDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
