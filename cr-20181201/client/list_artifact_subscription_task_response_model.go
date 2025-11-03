// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactSubscriptionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListArtifactSubscriptionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListArtifactSubscriptionTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListArtifactSubscriptionTaskResponseBody) *ListArtifactSubscriptionTaskResponse
	GetBody() *ListArtifactSubscriptionTaskResponseBody
}

type ListArtifactSubscriptionTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactSubscriptionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactSubscriptionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactSubscriptionTaskResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListArtifactSubscriptionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListArtifactSubscriptionTaskResponse) GetBody() *ListArtifactSubscriptionTaskResponseBody {
	return s.Body
}

func (s *ListArtifactSubscriptionTaskResponse) SetHeaders(v map[string]*string) *ListArtifactSubscriptionTaskResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactSubscriptionTaskResponse) SetStatusCode(v int32) *ListArtifactSubscriptionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponse) SetBody(v *ListArtifactSubscriptionTaskResponseBody) *ListArtifactSubscriptionTaskResponse {
	s.Body = v
	return s
}

func (s *ListArtifactSubscriptionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
