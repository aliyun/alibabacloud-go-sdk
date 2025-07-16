// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReportDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReportDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryReportDetailResponseBody) *QueryReportDetailResponse
	GetBody() *QueryReportDetailResponseBody
}

type QueryReportDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReportDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReportDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryReportDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReportDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReportDetailResponse) GetBody() *QueryReportDetailResponseBody {
	return s.Body
}

func (s *QueryReportDetailResponse) SetHeaders(v map[string]*string) *QueryReportDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryReportDetailResponse) SetStatusCode(v int32) *QueryReportDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReportDetailResponse) SetBody(v *QueryReportDetailResponseBody) *QueryReportDetailResponse {
	s.Body = v
	return s
}

func (s *QueryReportDetailResponse) Validate() error {
	return dara.Validate(s)
}
