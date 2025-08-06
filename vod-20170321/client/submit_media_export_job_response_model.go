// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMediaExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMediaExportJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMediaExportJobResponseBody) *SubmitMediaExportJobResponse
	GetBody() *SubmitMediaExportJobResponseBody
}

type SubmitMediaExportJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMediaExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMediaExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaExportJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMediaExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMediaExportJobResponse) GetBody() *SubmitMediaExportJobResponseBody {
	return s.Body
}

func (s *SubmitMediaExportJobResponse) SetHeaders(v map[string]*string) *SubmitMediaExportJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaExportJobResponse) SetStatusCode(v int32) *SubmitMediaExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMediaExportJobResponse) SetBody(v *SubmitMediaExportJobResponseBody) *SubmitMediaExportJobResponse {
	s.Body = v
	return s
}

func (s *SubmitMediaExportJobResponse) Validate() error {
	return dara.Validate(s)
}
