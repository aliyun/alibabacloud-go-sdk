// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDataExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendDataExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendDataExportJobResponse
	GetStatusCode() *int32
	SetBody(v *SuspendDataExportJobResponseBody) *SuspendDataExportJobResponse
	GetBody() *SuspendDataExportJobResponseBody
}

type SuspendDataExportJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendDataExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendDataExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendDataExportJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendDataExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendDataExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendDataExportJobResponse) GetBody() *SuspendDataExportJobResponseBody {
	return s.Body
}

func (s *SuspendDataExportJobResponse) SetHeaders(v map[string]*string) *SuspendDataExportJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendDataExportJobResponse) SetStatusCode(v int32) *SuspendDataExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendDataExportJobResponse) SetBody(v *SuspendDataExportJobResponseBody) *SuspendDataExportJobResponse {
	s.Body = v
	return s
}

func (s *SuspendDataExportJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
