// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnalysisExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAnalysisExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAnalysisExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAnalysisExportTaskResponseBody) *CreateAnalysisExportTaskResponse
	GetBody() *CreateAnalysisExportTaskResponseBody
}

type CreateAnalysisExportTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnalysisExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnalysisExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAnalysisExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAnalysisExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAnalysisExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAnalysisExportTaskResponse) GetBody() *CreateAnalysisExportTaskResponseBody {
	return s.Body
}

func (s *CreateAnalysisExportTaskResponse) SetHeaders(v map[string]*string) *CreateAnalysisExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAnalysisExportTaskResponse) SetStatusCode(v int32) *CreateAnalysisExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnalysisExportTaskResponse) SetBody(v *CreateAnalysisExportTaskResponseBody) *CreateAnalysisExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAnalysisExportTaskResponse) Validate() error {
	return dara.Validate(s)
}
