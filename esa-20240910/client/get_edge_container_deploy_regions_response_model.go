// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerDeployRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerDeployRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerDeployRegionsResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerDeployRegionsResponseBody) *GetEdgeContainerDeployRegionsResponse
	GetBody() *GetEdgeContainerDeployRegionsResponseBody
}

type GetEdgeContainerDeployRegionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerDeployRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerDeployRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerDeployRegionsResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerDeployRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerDeployRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerDeployRegionsResponse) GetBody() *GetEdgeContainerDeployRegionsResponseBody {
	return s.Body
}

func (s *GetEdgeContainerDeployRegionsResponse) SetHeaders(v map[string]*string) *GetEdgeContainerDeployRegionsResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerDeployRegionsResponse) SetStatusCode(v int32) *GetEdgeContainerDeployRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerDeployRegionsResponse) SetBody(v *GetEdgeContainerDeployRegionsResponseBody) *GetEdgeContainerDeployRegionsResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerDeployRegionsResponse) Validate() error {
	return dara.Validate(s)
}
