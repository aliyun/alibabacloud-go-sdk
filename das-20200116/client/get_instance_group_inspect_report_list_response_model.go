// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGroupInspectReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceGroupInspectReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceGroupInspectReportListResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceGroupInspectReportListResponseBody) *GetInstanceGroupInspectReportListResponse
	GetBody() *GetInstanceGroupInspectReportListResponseBody
}

type GetInstanceGroupInspectReportListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceGroupInspectReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceGroupInspectReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceGroupInspectReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceGroupInspectReportListResponse) GetBody() *GetInstanceGroupInspectReportListResponseBody {
	return s.Body
}

func (s *GetInstanceGroupInspectReportListResponse) SetHeaders(v map[string]*string) *GetInstanceGroupInspectReportListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceGroupInspectReportListResponse) SetStatusCode(v int32) *GetInstanceGroupInspectReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceGroupInspectReportListResponse) SetBody(v *GetInstanceGroupInspectReportListResponseBody) *GetInstanceGroupInspectReportListResponse {
	s.Body = v
	return s
}

func (s *GetInstanceGroupInspectReportListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
