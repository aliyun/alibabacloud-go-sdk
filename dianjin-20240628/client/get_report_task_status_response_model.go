// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReportTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReportTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetReportTaskStatusResponseBody) *GetReportTaskStatusResponse
	GetBody() *GetReportTaskStatusResponseBody
}

type GetReportTaskStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReportTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetReportTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReportTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReportTaskStatusResponse) GetBody() *GetReportTaskStatusResponseBody {
	return s.Body
}

func (s *GetReportTaskStatusResponse) SetHeaders(v map[string]*string) *GetReportTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetReportTaskStatusResponse) SetStatusCode(v int32) *GetReportTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportTaskStatusResponse) SetBody(v *GetReportTaskStatusResponseBody) *GetReportTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetReportTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
