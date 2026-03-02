// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryRepoMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLibraryRepoMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLibraryRepoMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetLibraryRepoMetricsResponseBody) *GetLibraryRepoMetricsResponse
	GetBody() *GetLibraryRepoMetricsResponseBody
}

type GetLibraryRepoMetricsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibraryRepoMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryRepoMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryRepoMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryRepoMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLibraryRepoMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLibraryRepoMetricsResponse) GetBody() *GetLibraryRepoMetricsResponseBody {
	return s.Body
}

func (s *GetLibraryRepoMetricsResponse) SetHeaders(v map[string]*string) *GetLibraryRepoMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryRepoMetricsResponse) SetStatusCode(v int32) *GetLibraryRepoMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryRepoMetricsResponse) SetBody(v *GetLibraryRepoMetricsResponseBody) *GetLibraryRepoMetricsResponse {
	s.Body = v
	return s
}

func (s *GetLibraryRepoMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
