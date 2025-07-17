// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDataExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDataExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDataExportJobResponse
	GetStatusCode() *int32
	SetBody(v *RestartDataExportJobResponseBody) *RestartDataExportJobResponse
	GetBody() *RestartDataExportJobResponseBody
}

type RestartDataExportJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDataExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDataExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDataExportJobResponse) GoString() string {
	return s.String()
}

func (s *RestartDataExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDataExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDataExportJobResponse) GetBody() *RestartDataExportJobResponseBody {
	return s.Body
}

func (s *RestartDataExportJobResponse) SetHeaders(v map[string]*string) *RestartDataExportJobResponse {
	s.Headers = v
	return s
}

func (s *RestartDataExportJobResponse) SetStatusCode(v int32) *RestartDataExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDataExportJobResponse) SetBody(v *RestartDataExportJobResponseBody) *RestartDataExportJobResponse {
	s.Body = v
	return s
}

func (s *RestartDataExportJobResponse) Validate() error {
	return dara.Validate(s)
}
