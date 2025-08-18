// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerStagingDeployStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeContainerStagingDeployStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeContainerStagingDeployStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeContainerStagingDeployStatusResponseBody) *GetEdgeContainerStagingDeployStatusResponse
	GetBody() *GetEdgeContainerStagingDeployStatusResponseBody
}

type GetEdgeContainerStagingDeployStatusResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeContainerStagingDeployStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeContainerStagingDeployStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerStagingDeployStatusResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerStagingDeployStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeContainerStagingDeployStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeContainerStagingDeployStatusResponse) GetBody() *GetEdgeContainerStagingDeployStatusResponseBody {
	return s.Body
}

func (s *GetEdgeContainerStagingDeployStatusResponse) SetHeaders(v map[string]*string) *GetEdgeContainerStagingDeployStatusResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponse) SetStatusCode(v int32) *GetEdgeContainerStagingDeployStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponse) SetBody(v *GetEdgeContainerStagingDeployStatusResponseBody) *GetEdgeContainerStagingDeployStatusResponse {
	s.Body = v
	return s
}

func (s *GetEdgeContainerStagingDeployStatusResponse) Validate() error {
	return dara.Validate(s)
}
