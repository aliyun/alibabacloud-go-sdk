// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchArtifactDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchArtifactDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchArtifactDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *ArtifactCode) *FetchArtifactDownloadUrlResponse
	GetBody() *ArtifactCode
}

type FetchArtifactDownloadUrlResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ArtifactCode      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchArtifactDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchArtifactDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *FetchArtifactDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchArtifactDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchArtifactDownloadUrlResponse) GetBody() *ArtifactCode {
	return s.Body
}

func (s *FetchArtifactDownloadUrlResponse) SetHeaders(v map[string]*string) *FetchArtifactDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *FetchArtifactDownloadUrlResponse) SetStatusCode(v int32) *FetchArtifactDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchArtifactDownloadUrlResponse) SetBody(v *ArtifactCode) *FetchArtifactDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *FetchArtifactDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}
