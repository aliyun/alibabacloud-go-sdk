// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretVersionStageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecretVersionStageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecretVersionStageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecretVersionStageResponseBody) *UpdateSecretVersionStageResponse
	GetBody() *UpdateSecretVersionStageResponseBody
}

type UpdateSecretVersionStageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecretVersionStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecretVersionStageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretVersionStageResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecretVersionStageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecretVersionStageResponse) GetBody() *UpdateSecretVersionStageResponseBody {
	return s.Body
}

func (s *UpdateSecretVersionStageResponse) SetHeaders(v map[string]*string) *UpdateSecretVersionStageResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretVersionStageResponse) SetStatusCode(v int32) *UpdateSecretVersionStageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecretVersionStageResponse) SetBody(v *UpdateSecretVersionStageResponseBody) *UpdateSecretVersionStageResponse {
	s.Body = v
	return s
}

func (s *UpdateSecretVersionStageResponse) Validate() error {
	return dara.Validate(s)
}
