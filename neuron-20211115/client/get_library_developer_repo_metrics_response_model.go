// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryDeveloperRepoMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLibraryDeveloperRepoMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLibraryDeveloperRepoMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetLibraryDeveloperRepoMetricsResponseBody) *GetLibraryDeveloperRepoMetricsResponse
	GetBody() *GetLibraryDeveloperRepoMetricsResponseBody
}

type GetLibraryDeveloperRepoMetricsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibraryDeveloperRepoMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryDeveloperRepoMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryDeveloperRepoMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryDeveloperRepoMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLibraryDeveloperRepoMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLibraryDeveloperRepoMetricsResponse) GetBody() *GetLibraryDeveloperRepoMetricsResponseBody {
	return s.Body
}

func (s *GetLibraryDeveloperRepoMetricsResponse) SetHeaders(v map[string]*string) *GetLibraryDeveloperRepoMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryDeveloperRepoMetricsResponse) SetStatusCode(v int32) *GetLibraryDeveloperRepoMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryDeveloperRepoMetricsResponse) SetBody(v *GetLibraryDeveloperRepoMetricsResponseBody) *GetLibraryDeveloperRepoMetricsResponse {
	s.Body = v
	return s
}

func (s *GetLibraryDeveloperRepoMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
