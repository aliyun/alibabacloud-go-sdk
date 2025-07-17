// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPromClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallPromClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallPromClusterResponse
	GetStatusCode() *int32
	SetBody(v *UninstallPromClusterResponseBody) *UninstallPromClusterResponse
	GetBody() *UninstallPromClusterResponseBody
}

type UninstallPromClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallPromClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallPromClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallPromClusterResponse) GoString() string {
	return s.String()
}

func (s *UninstallPromClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallPromClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallPromClusterResponse) GetBody() *UninstallPromClusterResponseBody {
	return s.Body
}

func (s *UninstallPromClusterResponse) SetHeaders(v map[string]*string) *UninstallPromClusterResponse {
	s.Headers = v
	return s
}

func (s *UninstallPromClusterResponse) SetStatusCode(v int32) *UninstallPromClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallPromClusterResponse) SetBody(v *UninstallPromClusterResponseBody) *UninstallPromClusterResponse {
	s.Body = v
	return s
}

func (s *UninstallPromClusterResponse) Validate() error {
	return dara.Validate(s)
}
