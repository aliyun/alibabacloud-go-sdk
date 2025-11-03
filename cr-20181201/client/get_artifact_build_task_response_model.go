// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactBuildTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactBuildTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactBuildTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetArtifactBuildTaskResponseBody) *GetArtifactBuildTaskResponse
	GetBody() *GetArtifactBuildTaskResponseBody
}

type GetArtifactBuildTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactBuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactBuildTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildTaskResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactBuildTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactBuildTaskResponse) GetBody() *GetArtifactBuildTaskResponseBody {
	return s.Body
}

func (s *GetArtifactBuildTaskResponse) SetHeaders(v map[string]*string) *GetArtifactBuildTaskResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactBuildTaskResponse) SetStatusCode(v int32) *GetArtifactBuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactBuildTaskResponse) SetBody(v *GetArtifactBuildTaskResponseBody) *GetArtifactBuildTaskResponse {
	s.Body = v
	return s
}

func (s *GetArtifactBuildTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
