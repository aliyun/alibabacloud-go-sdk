// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveWorkItemArtifactAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResolveWorkItemArtifactAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResolveWorkItemArtifactAccessResponse
	GetStatusCode() *int32
	SetBody(v *ResolveWorkItemArtifactAccessResponseBody) *ResolveWorkItemArtifactAccessResponse
	GetBody() *ResolveWorkItemArtifactAccessResponseBody
}

type ResolveWorkItemArtifactAccessResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResolveWorkItemArtifactAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResolveWorkItemArtifactAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s ResolveWorkItemArtifactAccessResponse) GoString() string {
	return s.String()
}

func (s *ResolveWorkItemArtifactAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResolveWorkItemArtifactAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResolveWorkItemArtifactAccessResponse) GetBody() *ResolveWorkItemArtifactAccessResponseBody {
	return s.Body
}

func (s *ResolveWorkItemArtifactAccessResponse) SetHeaders(v map[string]*string) *ResolveWorkItemArtifactAccessResponse {
	s.Headers = v
	return s
}

func (s *ResolveWorkItemArtifactAccessResponse) SetStatusCode(v int32) *ResolveWorkItemArtifactAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessResponse) SetBody(v *ResolveWorkItemArtifactAccessResponseBody) *ResolveWorkItemArtifactAccessResponse {
	s.Body = v
	return s
}

func (s *ResolveWorkItemArtifactAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
