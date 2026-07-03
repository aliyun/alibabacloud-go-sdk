// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAiModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAiModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateAiModelProviderResponseBody) *CreateAiModelProviderResponse
	GetBody() *CreateAiModelProviderResponseBody
}

type CreateAiModelProviderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAiModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAiModelProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateAiModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAiModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAiModelProviderResponse) GetBody() *CreateAiModelProviderResponseBody {
	return s.Body
}

func (s *CreateAiModelProviderResponse) SetHeaders(v map[string]*string) *CreateAiModelProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateAiModelProviderResponse) SetStatusCode(v int32) *CreateAiModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiModelProviderResponse) SetBody(v *CreateAiModelProviderResponseBody) *CreateAiModelProviderResponse {
	s.Body = v
	return s
}

func (s *CreateAiModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
