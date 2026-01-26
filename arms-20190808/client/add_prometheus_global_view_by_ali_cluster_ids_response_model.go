// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusGlobalViewByAliClusterIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPrometheusGlobalViewByAliClusterIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPrometheusGlobalViewByAliClusterIdsResponse
	GetStatusCode() *int32
	SetBody(v *AddPrometheusGlobalViewByAliClusterIdsResponseBody) *AddPrometheusGlobalViewByAliClusterIdsResponse
	GetBody() *AddPrometheusGlobalViewByAliClusterIdsResponseBody
}

type AddPrometheusGlobalViewByAliClusterIdsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPrometheusGlobalViewByAliClusterIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponse) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) GetBody() *AddPrometheusGlobalViewByAliClusterIdsResponseBody {
	return s.Body
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) SetHeaders(v map[string]*string) *AddPrometheusGlobalViewByAliClusterIdsResponse {
	s.Headers = v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) SetStatusCode(v int32) *AddPrometheusGlobalViewByAliClusterIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) SetBody(v *AddPrometheusGlobalViewByAliClusterIdsResponseBody) *AddPrometheusGlobalViewByAliClusterIdsResponse {
	s.Body = v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
