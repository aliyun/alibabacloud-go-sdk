// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactBuildLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListArtifactBuildLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListArtifactBuildLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListArtifactBuildLogsResponseBody) *ListArtifactBuildLogsResponse
	GetBody() *ListArtifactBuildLogsResponseBody
}

type ListArtifactBuildLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactBuildLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactBuildLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildLogsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListArtifactBuildLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListArtifactBuildLogsResponse) GetBody() *ListArtifactBuildLogsResponseBody {
	return s.Body
}

func (s *ListArtifactBuildLogsResponse) SetHeaders(v map[string]*string) *ListArtifactBuildLogsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactBuildLogsResponse) SetStatusCode(v int32) *ListArtifactBuildLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactBuildLogsResponse) SetBody(v *ListArtifactBuildLogsResponseBody) *ListArtifactBuildLogsResponse {
	s.Body = v
	return s
}

func (s *ListArtifactBuildLogsResponse) Validate() error {
	return dara.Validate(s)
}
