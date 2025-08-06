// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaExportJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaExportJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaExportJobsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaExportJobsResponseBody) *DeleteMediaExportJobsResponse
	GetBody() *DeleteMediaExportJobsResponseBody
}

type DeleteMediaExportJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaExportJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaExportJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaExportJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaExportJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaExportJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaExportJobsResponse) GetBody() *DeleteMediaExportJobsResponseBody {
	return s.Body
}

func (s *DeleteMediaExportJobsResponse) SetHeaders(v map[string]*string) *DeleteMediaExportJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaExportJobsResponse) SetStatusCode(v int32) *DeleteMediaExportJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaExportJobsResponse) SetBody(v *DeleteMediaExportJobsResponseBody) *DeleteMediaExportJobsResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaExportJobsResponse) Validate() error {
	return dara.Validate(s)
}
