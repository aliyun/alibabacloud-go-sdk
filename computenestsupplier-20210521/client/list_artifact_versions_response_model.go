// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListArtifactVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListArtifactVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListArtifactVersionsResponseBody) *ListArtifactVersionsResponse
	GetBody() *ListArtifactVersionsResponseBody
}

type ListArtifactVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListArtifactVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListArtifactVersionsResponse) GetBody() *ListArtifactVersionsResponseBody {
	return s.Body
}

func (s *ListArtifactVersionsResponse) SetHeaders(v map[string]*string) *ListArtifactVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactVersionsResponse) SetStatusCode(v int32) *ListArtifactVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactVersionsResponse) SetBody(v *ListArtifactVersionsResponseBody) *ListArtifactVersionsResponse {
	s.Body = v
	return s
}

func (s *ListArtifactVersionsResponse) Validate() error {
	return dara.Validate(s)
}
