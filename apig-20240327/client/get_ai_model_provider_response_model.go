// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiModelProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiModelProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetAiModelProviderResponseBody) *GetAiModelProviderResponse
	GetBody() *GetAiModelProviderResponseBody
}

type GetAiModelProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiModelProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiModelProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelProviderResponse) GoString() string {
	return s.String()
}

func (s *GetAiModelProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiModelProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiModelProviderResponse) GetBody() *GetAiModelProviderResponseBody {
	return s.Body
}

func (s *GetAiModelProviderResponse) SetHeaders(v map[string]*string) *GetAiModelProviderResponse {
	s.Headers = v
	return s
}

func (s *GetAiModelProviderResponse) SetStatusCode(v int32) *GetAiModelProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiModelProviderResponse) SetBody(v *GetAiModelProviderResponseBody) *GetAiModelProviderResponse {
	s.Body = v
	return s
}

func (s *GetAiModelProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
