// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaRefreshJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaRefreshJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaRefreshJobsResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaRefreshJobsResponseBody) *GetMediaRefreshJobsResponse
	GetBody() *GetMediaRefreshJobsResponseBody
}

type GetMediaRefreshJobsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaRefreshJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaRefreshJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaRefreshJobsResponse) GoString() string {
	return s.String()
}

func (s *GetMediaRefreshJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaRefreshJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaRefreshJobsResponse) GetBody() *GetMediaRefreshJobsResponseBody {
	return s.Body
}

func (s *GetMediaRefreshJobsResponse) SetHeaders(v map[string]*string) *GetMediaRefreshJobsResponse {
	s.Headers = v
	return s
}

func (s *GetMediaRefreshJobsResponse) SetStatusCode(v int32) *GetMediaRefreshJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaRefreshJobsResponse) SetBody(v *GetMediaRefreshJobsResponseBody) *GetMediaRefreshJobsResponse {
	s.Body = v
	return s
}

func (s *GetMediaRefreshJobsResponse) Validate() error {
	return dara.Validate(s)
}
