// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactRepositoryCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactRepositoryCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactRepositoryCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *GetArtifactRepositoryCredentialsResponseBody) *GetArtifactRepositoryCredentialsResponse
	GetBody() *GetArtifactRepositoryCredentialsResponseBody
}

type GetArtifactRepositoryCredentialsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactRepositoryCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactRepositoryCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactRepositoryCredentialsResponse) GetBody() *GetArtifactRepositoryCredentialsResponseBody {
	return s.Body
}

func (s *GetArtifactRepositoryCredentialsResponse) SetHeaders(v map[string]*string) *GetArtifactRepositoryCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) SetStatusCode(v int32) *GetArtifactRepositoryCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) SetBody(v *GetArtifactRepositoryCredentialsResponseBody) *GetArtifactRepositoryCredentialsResponse {
	s.Body = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) Validate() error {
	return dara.Validate(s)
}
