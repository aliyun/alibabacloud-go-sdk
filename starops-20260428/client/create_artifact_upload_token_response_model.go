// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactUploadTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateArtifactUploadTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateArtifactUploadTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateArtifactUploadTokenResponseBody) *CreateArtifactUploadTokenResponse
	GetBody() *CreateArtifactUploadTokenResponseBody
}

type CreateArtifactUploadTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactUploadTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactUploadTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateArtifactUploadTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateArtifactUploadTokenResponse) GetBody() *CreateArtifactUploadTokenResponseBody {
	return s.Body
}

func (s *CreateArtifactUploadTokenResponse) SetHeaders(v map[string]*string) *CreateArtifactUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactUploadTokenResponse) SetStatusCode(v int32) *CreateArtifactUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactUploadTokenResponse) SetBody(v *CreateArtifactUploadTokenResponseBody) *CreateArtifactUploadTokenResponse {
	s.Body = v
	return s
}

func (s *CreateArtifactUploadTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
