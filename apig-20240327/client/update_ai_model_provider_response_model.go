// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAiModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAiModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAiModelProviderResponseBody) *UpdateAiModelProviderResponse
	GetBody() *UpdateAiModelProviderResponseBody
}

type UpdateAiModelProviderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAiModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAiModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiModelProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateAiModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAiModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAiModelProviderResponse) GetBody() *UpdateAiModelProviderResponseBody {
	return s.Body
}

func (s *UpdateAiModelProviderResponse) SetHeaders(v map[string]*string) *UpdateAiModelProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateAiModelProviderResponse) SetStatusCode(v int32) *UpdateAiModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAiModelProviderResponse) SetBody(v *UpdateAiModelProviderResponseBody) *UpdateAiModelProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateAiModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
