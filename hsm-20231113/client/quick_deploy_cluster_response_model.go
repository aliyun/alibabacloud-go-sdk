// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickDeployClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuickDeployClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuickDeployClusterResponse
	GetStatusCode() *int32
	SetBody(v *QuickDeployClusterResponseBody) *QuickDeployClusterResponse
	GetBody() *QuickDeployClusterResponseBody
}

type QuickDeployClusterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuickDeployClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuickDeployClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s QuickDeployClusterResponse) GoString() string {
	return s.String()
}

func (s *QuickDeployClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuickDeployClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuickDeployClusterResponse) GetBody() *QuickDeployClusterResponseBody {
	return s.Body
}

func (s *QuickDeployClusterResponse) SetHeaders(v map[string]*string) *QuickDeployClusterResponse {
	s.Headers = v
	return s
}

func (s *QuickDeployClusterResponse) SetStatusCode(v int32) *QuickDeployClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *QuickDeployClusterResponse) SetBody(v *QuickDeployClusterResponseBody) *QuickDeployClusterResponse {
	s.Body = v
	return s
}

func (s *QuickDeployClusterResponse) Validate() error {
	return dara.Validate(s)
}
