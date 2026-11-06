// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNodeReportDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryNodeReportDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryNodeReportDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryNodeReportDataResponseBody) *QueryNodeReportDataResponse
	GetBody() *QueryNodeReportDataResponseBody
}

type QueryNodeReportDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNodeReportDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNodeReportDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryNodeReportDataResponse) GoString() string {
	return s.String()
}

func (s *QueryNodeReportDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryNodeReportDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryNodeReportDataResponse) GetBody() *QueryNodeReportDataResponseBody {
	return s.Body
}

func (s *QueryNodeReportDataResponse) SetHeaders(v map[string]*string) *QueryNodeReportDataResponse {
	s.Headers = v
	return s
}

func (s *QueryNodeReportDataResponse) SetStatusCode(v int32) *QueryNodeReportDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNodeReportDataResponse) SetBody(v *QueryNodeReportDataResponseBody) *QueryNodeReportDataResponse {
	s.Body = v
	return s
}

func (s *QueryNodeReportDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
