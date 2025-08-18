// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildEdgeContainerAppStagingEnvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebuildEdgeContainerAppStagingEnvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebuildEdgeContainerAppStagingEnvResponse
	GetStatusCode() *int32
	SetBody(v *RebuildEdgeContainerAppStagingEnvResponseBody) *RebuildEdgeContainerAppStagingEnvResponse
	GetBody() *RebuildEdgeContainerAppStagingEnvResponseBody
}

type RebuildEdgeContainerAppStagingEnvResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebuildEdgeContainerAppStagingEnvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebuildEdgeContainerAppStagingEnvResponse) String() string {
	return dara.Prettify(s)
}

func (s RebuildEdgeContainerAppStagingEnvResponse) GoString() string {
	return s.String()
}

func (s *RebuildEdgeContainerAppStagingEnvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebuildEdgeContainerAppStagingEnvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebuildEdgeContainerAppStagingEnvResponse) GetBody() *RebuildEdgeContainerAppStagingEnvResponseBody {
	return s.Body
}

func (s *RebuildEdgeContainerAppStagingEnvResponse) SetHeaders(v map[string]*string) *RebuildEdgeContainerAppStagingEnvResponse {
	s.Headers = v
	return s
}

func (s *RebuildEdgeContainerAppStagingEnvResponse) SetStatusCode(v int32) *RebuildEdgeContainerAppStagingEnvResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildEdgeContainerAppStagingEnvResponse) SetBody(v *RebuildEdgeContainerAppStagingEnvResponseBody) *RebuildEdgeContainerAppStagingEnvResponse {
	s.Body = v
	return s
}

func (s *RebuildEdgeContainerAppStagingEnvResponse) Validate() error {
	return dara.Validate(s)
}
