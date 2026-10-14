// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListGpuClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetListGpuClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetListGpuClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetListGpuClusterResponseBody) *GetListGpuClusterResponse
	GetBody() *GetListGpuClusterResponseBody
}

type GetListGpuClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetListGpuClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetListGpuClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetListGpuClusterResponse) GoString() string {
	return s.String()
}

func (s *GetListGpuClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetListGpuClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetListGpuClusterResponse) GetBody() *GetListGpuClusterResponseBody {
	return s.Body
}

func (s *GetListGpuClusterResponse) SetHeaders(v map[string]*string) *GetListGpuClusterResponse {
	s.Headers = v
	return s
}

func (s *GetListGpuClusterResponse) SetStatusCode(v int32) *GetListGpuClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListGpuClusterResponse) SetBody(v *GetListGpuClusterResponseBody) *GetListGpuClusterResponse {
	s.Body = v
	return s
}

func (s *GetListGpuClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
