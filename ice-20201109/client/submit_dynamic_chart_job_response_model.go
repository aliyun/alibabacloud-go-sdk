// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicChartJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDynamicChartJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDynamicChartJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDynamicChartJobResponseBody) *SubmitDynamicChartJobResponse
	GetBody() *SubmitDynamicChartJobResponseBody
}

type SubmitDynamicChartJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDynamicChartJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDynamicChartJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicChartJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDynamicChartJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDynamicChartJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDynamicChartJobResponse) GetBody() *SubmitDynamicChartJobResponseBody {
	return s.Body
}

func (s *SubmitDynamicChartJobResponse) SetHeaders(v map[string]*string) *SubmitDynamicChartJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDynamicChartJobResponse) SetStatusCode(v int32) *SubmitDynamicChartJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDynamicChartJobResponse) SetBody(v *SubmitDynamicChartJobResponseBody) *SubmitDynamicChartJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDynamicChartJobResponse) Validate() error {
	return dara.Validate(s)
}
