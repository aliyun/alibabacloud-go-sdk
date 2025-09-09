// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListArtifactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListArtifactsResponse
	GetStatusCode() *int32
	SetBody(v *ListArtifactsResponseBody) *ListArtifactsResponse
	GetBody() *ListArtifactsResponseBody
}

type ListArtifactsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListArtifactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListArtifactsResponse) GetBody() *ListArtifactsResponseBody {
	return s.Body
}

func (s *ListArtifactsResponse) SetHeaders(v map[string]*string) *ListArtifactsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactsResponse) SetStatusCode(v int32) *ListArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactsResponse) SetBody(v *ListArtifactsResponseBody) *ListArtifactsResponse {
	s.Body = v
	return s
}

func (s *ListArtifactsResponse) Validate() error {
	return dara.Validate(s)
}
