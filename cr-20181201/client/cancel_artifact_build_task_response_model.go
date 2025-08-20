// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelArtifactBuildTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelArtifactBuildTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelArtifactBuildTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelArtifactBuildTaskResponseBody) *CancelArtifactBuildTaskResponse
	GetBody() *CancelArtifactBuildTaskResponseBody
}

type CancelArtifactBuildTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelArtifactBuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelArtifactBuildTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelArtifactBuildTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelArtifactBuildTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelArtifactBuildTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelArtifactBuildTaskResponse) GetBody() *CancelArtifactBuildTaskResponseBody {
	return s.Body
}

func (s *CancelArtifactBuildTaskResponse) SetHeaders(v map[string]*string) *CancelArtifactBuildTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelArtifactBuildTaskResponse) SetStatusCode(v int32) *CancelArtifactBuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelArtifactBuildTaskResponse) SetBody(v *CancelArtifactBuildTaskResponseBody) *CancelArtifactBuildTaskResponse {
	s.Body = v
	return s
}

func (s *CancelArtifactBuildTaskResponse) Validate() error {
	return dara.Validate(s)
}
