// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactSubscriptionTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactSubscriptionTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetArtifactSubscriptionTaskResultResponseBody) *GetArtifactSubscriptionTaskResultResponse
	GetBody() *GetArtifactSubscriptionTaskResultResponseBody
}

type GetArtifactSubscriptionTaskResultResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactSubscriptionTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactSubscriptionTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactSubscriptionTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactSubscriptionTaskResultResponse) GetBody() *GetArtifactSubscriptionTaskResultResponseBody {
	return s.Body
}

func (s *GetArtifactSubscriptionTaskResultResponse) SetHeaders(v map[string]*string) *GetArtifactSubscriptionTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponse) SetStatusCode(v int32) *GetArtifactSubscriptionTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponse) SetBody(v *GetArtifactSubscriptionTaskResultResponseBody) *GetArtifactSubscriptionTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponse) Validate() error {
	return dara.Validate(s)
}
