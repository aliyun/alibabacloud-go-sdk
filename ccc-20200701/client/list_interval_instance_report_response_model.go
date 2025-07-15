// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalInstanceReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntervalInstanceReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntervalInstanceReportResponse
	GetStatusCode() *int32
	SetBody(v *ListIntervalInstanceReportResponseBody) *ListIntervalInstanceReportResponse
	GetBody() *ListIntervalInstanceReportResponseBody
}

type ListIntervalInstanceReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntervalInstanceReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntervalInstanceReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalInstanceReportResponse) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntervalInstanceReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntervalInstanceReportResponse) GetBody() *ListIntervalInstanceReportResponseBody {
	return s.Body
}

func (s *ListIntervalInstanceReportResponse) SetHeaders(v map[string]*string) *ListIntervalInstanceReportResponse {
	s.Headers = v
	return s
}

func (s *ListIntervalInstanceReportResponse) SetStatusCode(v int32) *ListIntervalInstanceReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntervalInstanceReportResponse) SetBody(v *ListIntervalInstanceReportResponseBody) *ListIntervalInstanceReportResponse {
	s.Body = v
	return s
}

func (s *ListIntervalInstanceReportResponse) Validate() error {
	return dara.Validate(s)
}
