// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateArtifactResponse
	GetStatusCode() *int32
	SetBody(v *UpdateArtifactResponseBody) *UpdateArtifactResponse
	GetBody() *UpdateArtifactResponseBody
}

type UpdateArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactResponse) GoString() string {
	return s.String()
}

func (s *UpdateArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateArtifactResponse) GetBody() *UpdateArtifactResponseBody {
	return s.Body
}

func (s *UpdateArtifactResponse) SetHeaders(v map[string]*string) *UpdateArtifactResponse {
	s.Headers = v
	return s
}

func (s *UpdateArtifactResponse) SetStatusCode(v int32) *UpdateArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateArtifactResponse) SetBody(v *UpdateArtifactResponseBody) *UpdateArtifactResponse {
	s.Body = v
	return s
}

func (s *UpdateArtifactResponse) Validate() error {
	return dara.Validate(s)
}
