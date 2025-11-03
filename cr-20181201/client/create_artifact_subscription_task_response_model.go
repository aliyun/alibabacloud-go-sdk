// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactSubscriptionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateArtifactSubscriptionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateArtifactSubscriptionTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateArtifactSubscriptionTaskResponseBody) *CreateArtifactSubscriptionTaskResponse
	GetBody() *CreateArtifactSubscriptionTaskResponseBody
}

type CreateArtifactSubscriptionTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactSubscriptionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactSubscriptionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactSubscriptionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateArtifactSubscriptionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateArtifactSubscriptionTaskResponse) GetBody() *CreateArtifactSubscriptionTaskResponseBody {
	return s.Body
}

func (s *CreateArtifactSubscriptionTaskResponse) SetHeaders(v map[string]*string) *CreateArtifactSubscriptionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponse) SetStatusCode(v int32) *CreateArtifactSubscriptionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponse) SetBody(v *CreateArtifactSubscriptionTaskResponseBody) *CreateArtifactSubscriptionTaskResponse {
	s.Body = v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
