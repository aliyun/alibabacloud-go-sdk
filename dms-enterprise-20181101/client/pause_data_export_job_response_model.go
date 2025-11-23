// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseDataExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseDataExportJobResponse
	GetStatusCode() *int32
	SetBody(v *PauseDataExportJobResponseBody) *PauseDataExportJobResponse
	GetBody() *PauseDataExportJobResponseBody
}

type PauseDataExportJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseDataExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseDataExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseDataExportJobResponse) GoString() string {
	return s.String()
}

func (s *PauseDataExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseDataExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseDataExportJobResponse) GetBody() *PauseDataExportJobResponseBody {
	return s.Body
}

func (s *PauseDataExportJobResponse) SetHeaders(v map[string]*string) *PauseDataExportJobResponse {
	s.Headers = v
	return s
}

func (s *PauseDataExportJobResponse) SetStatusCode(v int32) *PauseDataExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseDataExportJobResponse) SetBody(v *PauseDataExportJobResponseBody) *PauseDataExportJobResponse {
	s.Body = v
	return s
}

func (s *PauseDataExportJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
