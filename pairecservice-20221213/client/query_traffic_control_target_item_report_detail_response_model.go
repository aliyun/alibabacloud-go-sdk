// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTargetItemReportDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTrafficControlTargetItemReportDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTrafficControlTargetItemReportDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryTrafficControlTargetItemReportDetailResponseBody) *QueryTrafficControlTargetItemReportDetailResponse
	GetBody() *QueryTrafficControlTargetItemReportDetailResponseBody
}

type QueryTrafficControlTargetItemReportDetailResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTrafficControlTargetItemReportDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTrafficControlTargetItemReportDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTargetItemReportDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) GetBody() *QueryTrafficControlTargetItemReportDetailResponseBody {
	return s.Body
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) SetHeaders(v map[string]*string) *QueryTrafficControlTargetItemReportDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) SetStatusCode(v int32) *QueryTrafficControlTargetItemReportDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) SetBody(v *QueryTrafficControlTargetItemReportDetailResponseBody) *QueryTrafficControlTargetItemReportDetailResponse {
	s.Body = v
	return s
}

func (s *QueryTrafficControlTargetItemReportDetailResponse) Validate() error {
	return dara.Validate(s)
}
