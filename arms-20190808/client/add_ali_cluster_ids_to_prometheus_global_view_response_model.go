// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAliClusterIdsToPrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAliClusterIdsToPrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAliClusterIdsToPrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *AddAliClusterIdsToPrometheusGlobalViewResponseBody) *AddAliClusterIdsToPrometheusGlobalViewResponse
	GetBody() *AddAliClusterIdsToPrometheusGlobalViewResponseBody
}

type AddAliClusterIdsToPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAliClusterIdsToPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) GetBody() *AddAliClusterIdsToPrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *AddAliClusterIdsToPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) SetStatusCode(v int32) *AddAliClusterIdsToPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) SetBody(v *AddAliClusterIdsToPrometheusGlobalViewResponseBody) *AddAliClusterIdsToPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) Validate() error {
	return dara.Validate(s)
}
