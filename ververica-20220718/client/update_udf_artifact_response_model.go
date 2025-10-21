// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUdfArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUdfArtifactResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUdfArtifactResponseBody) *UpdateUdfArtifactResponse
	GetBody() *UpdateUdfArtifactResponseBody
}

type UpdateUdfArtifactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUdfArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUdfArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfArtifactResponse) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUdfArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUdfArtifactResponse) GetBody() *UpdateUdfArtifactResponseBody {
	return s.Body
}

func (s *UpdateUdfArtifactResponse) SetHeaders(v map[string]*string) *UpdateUdfArtifactResponse {
	s.Headers = v
	return s
}

func (s *UpdateUdfArtifactResponse) SetStatusCode(v int32) *UpdateUdfArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUdfArtifactResponse) SetBody(v *UpdateUdfArtifactResponseBody) *UpdateUdfArtifactResponse {
	s.Body = v
	return s
}

func (s *UpdateUdfArtifactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
