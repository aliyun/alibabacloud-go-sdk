// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDtsJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDtsJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDtsJobsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDtsJobsResponseBody) *DeleteDtsJobsResponse
	GetBody() *DeleteDtsJobsResponseBody
}

type DeleteDtsJobsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDtsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDtsJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDtsJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDtsJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDtsJobsResponse) GetBody() *DeleteDtsJobsResponseBody {
	return s.Body
}

func (s *DeleteDtsJobsResponse) SetHeaders(v map[string]*string) *DeleteDtsJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDtsJobsResponse) SetStatusCode(v int32) *DeleteDtsJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDtsJobsResponse) SetBody(v *DeleteDtsJobsResponseBody) *DeleteDtsJobsResponse {
	s.Body = v
	return s
}

func (s *DeleteDtsJobsResponse) Validate() error {
	return dara.Validate(s)
}
