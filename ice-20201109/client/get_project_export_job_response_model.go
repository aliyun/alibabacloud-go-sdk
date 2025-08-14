// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectExportJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectExportJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectExportJobResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectExportJobResponseBody) *GetProjectExportJobResponse
	GetBody() *GetProjectExportJobResponseBody
}

type GetProjectExportJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectExportJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectExportJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectExportJobResponse) GoString() string {
	return s.String()
}

func (s *GetProjectExportJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectExportJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectExportJobResponse) GetBody() *GetProjectExportJobResponseBody {
	return s.Body
}

func (s *GetProjectExportJobResponse) SetHeaders(v map[string]*string) *GetProjectExportJobResponse {
	s.Headers = v
	return s
}

func (s *GetProjectExportJobResponse) SetStatusCode(v int32) *GetProjectExportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectExportJobResponse) SetBody(v *GetProjectExportJobResponseBody) *GetProjectExportJobResponse {
	s.Body = v
	return s
}

func (s *GetProjectExportJobResponse) Validate() error {
	return dara.Validate(s)
}
