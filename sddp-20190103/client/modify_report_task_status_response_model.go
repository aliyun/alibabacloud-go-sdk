// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReportTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyReportTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyReportTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyReportTaskStatusResponseBody) *ModifyReportTaskStatusResponse
	GetBody() *ModifyReportTaskStatusResponseBody
}

type ModifyReportTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReportTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReportTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyReportTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyReportTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyReportTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyReportTaskStatusResponse) GetBody() *ModifyReportTaskStatusResponseBody {
	return s.Body
}

func (s *ModifyReportTaskStatusResponse) SetHeaders(v map[string]*string) *ModifyReportTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyReportTaskStatusResponse) SetStatusCode(v int32) *ModifyReportTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReportTaskStatusResponse) SetBody(v *ModifyReportTaskStatusResponseBody) *ModifyReportTaskStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyReportTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
