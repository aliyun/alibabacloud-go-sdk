// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterArgsOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGpuClusterArgsOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGpuClusterArgsOptionsResponse
	GetStatusCode() *int32
	SetBody(v *GetGpuClusterArgsOptionsResponseBody) *GetGpuClusterArgsOptionsResponse
	GetBody() *GetGpuClusterArgsOptionsResponseBody
}

type GetGpuClusterArgsOptionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGpuClusterArgsOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGpuClusterArgsOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterArgsOptionsResponse) GoString() string {
	return s.String()
}

func (s *GetGpuClusterArgsOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGpuClusterArgsOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGpuClusterArgsOptionsResponse) GetBody() *GetGpuClusterArgsOptionsResponseBody {
	return s.Body
}

func (s *GetGpuClusterArgsOptionsResponse) SetHeaders(v map[string]*string) *GetGpuClusterArgsOptionsResponse {
	s.Headers = v
	return s
}

func (s *GetGpuClusterArgsOptionsResponse) SetStatusCode(v int32) *GetGpuClusterArgsOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGpuClusterArgsOptionsResponse) SetBody(v *GetGpuClusterArgsOptionsResponseBody) *GetGpuClusterArgsOptionsResponse {
	s.Body = v
	return s
}

func (s *GetGpuClusterArgsOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
