// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMediaExportJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelMediaExportJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelMediaExportJobsResponse
	GetStatusCode() *int32
	SetBody(v *CancelMediaExportJobsResponseBody) *CancelMediaExportJobsResponse
	GetBody() *CancelMediaExportJobsResponseBody
}

type CancelMediaExportJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelMediaExportJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelMediaExportJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelMediaExportJobsResponse) GoString() string {
	return s.String()
}

func (s *CancelMediaExportJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelMediaExportJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelMediaExportJobsResponse) GetBody() *CancelMediaExportJobsResponseBody {
	return s.Body
}

func (s *CancelMediaExportJobsResponse) SetHeaders(v map[string]*string) *CancelMediaExportJobsResponse {
	s.Headers = v
	return s
}

func (s *CancelMediaExportJobsResponse) SetStatusCode(v int32) *CancelMediaExportJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelMediaExportJobsResponse) SetBody(v *CancelMediaExportJobsResponseBody) *CancelMediaExportJobsResponse {
	s.Body = v
	return s
}

func (s *CancelMediaExportJobsResponse) Validate() error {
	return dara.Validate(s)
}
