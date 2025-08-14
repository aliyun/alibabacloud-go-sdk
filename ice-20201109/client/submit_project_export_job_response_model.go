// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProjectExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitProjectExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitProjectExportJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitProjectExportJobResponseBody) *SubmitProjectExportJobResponse
	GetBody() *SubmitProjectExportJobResponseBody
}

type SubmitProjectExportJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitProjectExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitProjectExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectExportJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitProjectExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitProjectExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitProjectExportJobResponse) GetBody() *SubmitProjectExportJobResponseBody {
	return s.Body
}

func (s *SubmitProjectExportJobResponse) SetHeaders(v map[string]*string) *SubmitProjectExportJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitProjectExportJobResponse) SetStatusCode(v int32) *SubmitProjectExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitProjectExportJobResponse) SetBody(v *SubmitProjectExportJobResponseBody) *SubmitProjectExportJobResponse {
	s.Body = v
	return s
}

func (s *SubmitProjectExportJobResponse) Validate() error {
	return dara.Validate(s)
}
