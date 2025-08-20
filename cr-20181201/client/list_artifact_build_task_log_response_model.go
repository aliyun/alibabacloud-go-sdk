// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactBuildTaskLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListArtifactBuildTaskLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListArtifactBuildTaskLogResponse
	GetStatusCode() *int32
	SetBody(v *ListArtifactBuildTaskLogResponseBody) *ListArtifactBuildTaskLogResponse
	GetBody() *ListArtifactBuildTaskLogResponseBody
}

type ListArtifactBuildTaskLogResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactBuildTaskLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactBuildTaskLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildTaskLogResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildTaskLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListArtifactBuildTaskLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListArtifactBuildTaskLogResponse) GetBody() *ListArtifactBuildTaskLogResponseBody {
	return s.Body
}

func (s *ListArtifactBuildTaskLogResponse) SetHeaders(v map[string]*string) *ListArtifactBuildTaskLogResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactBuildTaskLogResponse) SetStatusCode(v int32) *ListArtifactBuildTaskLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponse) SetBody(v *ListArtifactBuildTaskLogResponseBody) *ListArtifactBuildTaskLogResponse {
	s.Body = v
	return s
}

func (s *ListArtifactBuildTaskLogResponse) Validate() error {
	return dara.Validate(s)
}
