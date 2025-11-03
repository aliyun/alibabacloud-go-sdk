// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactSubscriptionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactSubscriptionTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetArtifactSubscriptionTaskResponseBody) *GetArtifactSubscriptionTaskResponse
	GetBody() *GetArtifactSubscriptionTaskResponseBody
}

type GetArtifactSubscriptionTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactSubscriptionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactSubscriptionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactSubscriptionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactSubscriptionTaskResponse) GetBody() *GetArtifactSubscriptionTaskResponseBody {
	return s.Body
}

func (s *GetArtifactSubscriptionTaskResponse) SetHeaders(v map[string]*string) *GetArtifactSubscriptionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactSubscriptionTaskResponse) SetStatusCode(v int32) *GetArtifactSubscriptionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponse) SetBody(v *GetArtifactSubscriptionTaskResponseBody) *GetArtifactSubscriptionTaskResponse {
	s.Body = v
	return s
}

func (s *GetArtifactSubscriptionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
