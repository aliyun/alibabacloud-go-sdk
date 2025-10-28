// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetK8sClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetK8sClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetK8sClusterResponseBody) *GetK8sClusterResponse
	GetBody() *GetK8sClusterResponseBody
}

type GetK8sClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetK8sClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetK8sClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetK8sClusterResponse) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetK8sClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetK8sClusterResponse) GetBody() *GetK8sClusterResponseBody {
	return s.Body
}

func (s *GetK8sClusterResponse) SetHeaders(v map[string]*string) *GetK8sClusterResponse {
	s.Headers = v
	return s
}

func (s *GetK8sClusterResponse) SetStatusCode(v int32) *GetK8sClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetK8sClusterResponse) SetBody(v *GetK8sClusterResponseBody) *GetK8sClusterResponse {
	s.Body = v
	return s
}

func (s *GetK8sClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
