// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitResultExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitResultExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitResultExportJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitResultExportJobResponseBody) *SubmitResultExportJobResponse
	GetBody() *SubmitResultExportJobResponseBody
}

type SubmitResultExportJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitResultExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitResultExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitResultExportJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitResultExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitResultExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitResultExportJobResponse) GetBody() *SubmitResultExportJobResponseBody {
	return s.Body
}

func (s *SubmitResultExportJobResponse) SetHeaders(v map[string]*string) *SubmitResultExportJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitResultExportJobResponse) SetStatusCode(v int32) *SubmitResultExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitResultExportJobResponse) SetBody(v *SubmitResultExportJobResponseBody) *SubmitResultExportJobResponse {
	s.Body = v
	return s
}

func (s *SubmitResultExportJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
