// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUdfArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUdfArtifactResponse
	GetStatusCode() *int32
	SetBody(v *CreateUdfArtifactResponseBody) *CreateUdfArtifactResponse
	GetBody() *CreateUdfArtifactResponseBody
}

type CreateUdfArtifactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUdfArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUdfArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfArtifactResponse) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUdfArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUdfArtifactResponse) GetBody() *CreateUdfArtifactResponseBody {
	return s.Body
}

func (s *CreateUdfArtifactResponse) SetHeaders(v map[string]*string) *CreateUdfArtifactResponse {
	s.Headers = v
	return s
}

func (s *CreateUdfArtifactResponse) SetStatusCode(v int32) *CreateUdfArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUdfArtifactResponse) SetBody(v *CreateUdfArtifactResponseBody) *CreateUdfArtifactResponse {
	s.Body = v
	return s
}

func (s *CreateUdfArtifactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
