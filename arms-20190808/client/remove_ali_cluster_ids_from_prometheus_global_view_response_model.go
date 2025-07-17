// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAliClusterIdsFromPrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse
	GetBody() *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody
}

type RemoveAliClusterIdsFromPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) GetBody() *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) SetStatusCode(v int32) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) SetBody(v *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) Validate() error {
	return dara.Validate(s)
}
