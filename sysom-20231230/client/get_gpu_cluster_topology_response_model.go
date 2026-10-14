// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGpuClusterTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGpuClusterTopologyResponse
	GetStatusCode() *int32
	SetBody(v *GetGpuClusterTopologyResponseBody) *GetGpuClusterTopologyResponse
	GetBody() *GetGpuClusterTopologyResponseBody
}

type GetGpuClusterTopologyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGpuClusterTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGpuClusterTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterTopologyResponse) GoString() string {
	return s.String()
}

func (s *GetGpuClusterTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGpuClusterTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGpuClusterTopologyResponse) GetBody() *GetGpuClusterTopologyResponseBody {
	return s.Body
}

func (s *GetGpuClusterTopologyResponse) SetHeaders(v map[string]*string) *GetGpuClusterTopologyResponse {
	s.Headers = v
	return s
}

func (s *GetGpuClusterTopologyResponse) SetStatusCode(v int32) *GetGpuClusterTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGpuClusterTopologyResponse) SetBody(v *GetGpuClusterTopologyResponseBody) *GetGpuClusterTopologyResponse {
	s.Body = v
	return s
}

func (s *GetGpuClusterTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
