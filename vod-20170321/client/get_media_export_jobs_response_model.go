// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaExportJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaExportJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaExportJobsResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaExportJobsResponseBody) *GetMediaExportJobsResponse
	GetBody() *GetMediaExportJobsResponseBody
}

type GetMediaExportJobsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaExportJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaExportJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaExportJobsResponse) GoString() string {
	return s.String()
}

func (s *GetMediaExportJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaExportJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaExportJobsResponse) GetBody() *GetMediaExportJobsResponseBody {
	return s.Body
}

func (s *GetMediaExportJobsResponse) SetHeaders(v map[string]*string) *GetMediaExportJobsResponse {
	s.Headers = v
	return s
}

func (s *GetMediaExportJobsResponse) SetStatusCode(v int32) *GetMediaExportJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaExportJobsResponse) SetBody(v *GetMediaExportJobsResponseBody) *GetMediaExportJobsResponse {
	s.Body = v
	return s
}

func (s *GetMediaExportJobsResponse) Validate() error {
	return dara.Validate(s)
}
