// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportK8sClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportK8sClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportK8sClusterResponse
	GetStatusCode() *int32
	SetBody(v *ImportK8sClusterResponseBody) *ImportK8sClusterResponse
	GetBody() *ImportK8sClusterResponseBody
}

type ImportK8sClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportK8sClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportK8sClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportK8sClusterResponse) GoString() string {
	return s.String()
}

func (s *ImportK8sClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportK8sClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportK8sClusterResponse) GetBody() *ImportK8sClusterResponseBody {
	return s.Body
}

func (s *ImportK8sClusterResponse) SetHeaders(v map[string]*string) *ImportK8sClusterResponse {
	s.Headers = v
	return s
}

func (s *ImportK8sClusterResponse) SetStatusCode(v int32) *ImportK8sClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportK8sClusterResponse) SetBody(v *ImportK8sClusterResponseBody) *ImportK8sClusterResponse {
	s.Body = v
	return s
}

func (s *ImportK8sClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
