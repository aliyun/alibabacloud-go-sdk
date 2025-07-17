// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSourcesFromPrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSourcesFromPrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSourcesFromPrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSourcesFromPrometheusGlobalViewResponseBody) *RemoveSourcesFromPrometheusGlobalViewResponse
	GetBody() *RemoveSourcesFromPrometheusGlobalViewResponseBody
}

type RemoveSourcesFromPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSourcesFromPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSourcesFromPrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSourcesFromPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) GetBody() *RemoveSourcesFromPrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *RemoveSourcesFromPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) SetStatusCode(v int32) *RemoveSourcesFromPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) SetBody(v *RemoveSourcesFromPrometheusGlobalViewResponseBody) *RemoveSourcesFromPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) Validate() error {
	return dara.Validate(s)
}
