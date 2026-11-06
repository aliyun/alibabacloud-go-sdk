// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppInstanceGroupReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAppInstanceGroupReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAppInstanceGroupReportListResponse
	GetStatusCode() *int32
	SetBody(v *QueryAppInstanceGroupReportListResponseBody) *QueryAppInstanceGroupReportListResponse
	GetBody() *QueryAppInstanceGroupReportListResponseBody
}

type QueryAppInstanceGroupReportListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppInstanceGroupReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppInstanceGroupReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAppInstanceGroupReportListResponse) GoString() string {
	return s.String()
}

func (s *QueryAppInstanceGroupReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAppInstanceGroupReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAppInstanceGroupReportListResponse) GetBody() *QueryAppInstanceGroupReportListResponseBody {
	return s.Body
}

func (s *QueryAppInstanceGroupReportListResponse) SetHeaders(v map[string]*string) *QueryAppInstanceGroupReportListResponse {
	s.Headers = v
	return s
}

func (s *QueryAppInstanceGroupReportListResponse) SetStatusCode(v int32) *QueryAppInstanceGroupReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppInstanceGroupReportListResponse) SetBody(v *QueryAppInstanceGroupReportListResponseBody) *QueryAppInstanceGroupReportListResponse {
	s.Body = v
	return s
}

func (s *QueryAppInstanceGroupReportListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
