// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactRisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListArtifactRisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListArtifactRisksResponse
	GetStatusCode() *int32
	SetBody(v *ListArtifactRisksResponseBody) *ListArtifactRisksResponse
	GetBody() *ListArtifactRisksResponseBody
}

type ListArtifactRisksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactRisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactRisksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactRisksResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactRisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListArtifactRisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListArtifactRisksResponse) GetBody() *ListArtifactRisksResponseBody {
	return s.Body
}

func (s *ListArtifactRisksResponse) SetHeaders(v map[string]*string) *ListArtifactRisksResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactRisksResponse) SetStatusCode(v int32) *ListArtifactRisksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactRisksResponse) SetBody(v *ListArtifactRisksResponseBody) *ListArtifactRisksResponse {
	s.Body = v
	return s
}

func (s *ListArtifactRisksResponse) Validate() error {
	return dara.Validate(s)
}
