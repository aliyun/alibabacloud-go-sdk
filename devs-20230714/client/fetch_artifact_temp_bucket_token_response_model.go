// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchArtifactTempBucketTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchArtifactTempBucketTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchArtifactTempBucketTokenResponse
	GetStatusCode() *int32
	SetBody(v *ArtifactTempBucketToken) *FetchArtifactTempBucketTokenResponse
	GetBody() *ArtifactTempBucketToken
}

type FetchArtifactTempBucketTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ArtifactTempBucketToken `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchArtifactTempBucketTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchArtifactTempBucketTokenResponse) GoString() string {
	return s.String()
}

func (s *FetchArtifactTempBucketTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchArtifactTempBucketTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchArtifactTempBucketTokenResponse) GetBody() *ArtifactTempBucketToken {
	return s.Body
}

func (s *FetchArtifactTempBucketTokenResponse) SetHeaders(v map[string]*string) *FetchArtifactTempBucketTokenResponse {
	s.Headers = v
	return s
}

func (s *FetchArtifactTempBucketTokenResponse) SetStatusCode(v int32) *FetchArtifactTempBucketTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchArtifactTempBucketTokenResponse) SetBody(v *ArtifactTempBucketToken) *FetchArtifactTempBucketTokenResponse {
	s.Body = v
	return s
}

func (s *FetchArtifactTempBucketTokenResponse) Validate() error {
	return dara.Validate(s)
}
