// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetArtifactDownloadUrlResponseBody) *GetArtifactDownloadUrlResponse
	GetBody() *GetArtifactDownloadUrlResponseBody
}

type GetArtifactDownloadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactDownloadUrlResponse) GetBody() *GetArtifactDownloadUrlResponseBody {
	return s.Body
}

func (s *GetArtifactDownloadUrlResponse) SetHeaders(v map[string]*string) *GetArtifactDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactDownloadUrlResponse) SetStatusCode(v int32) *GetArtifactDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactDownloadUrlResponse) SetBody(v *GetArtifactDownloadUrlResponseBody) *GetArtifactDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetArtifactDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
