// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableLLMModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAvailableLLMModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAvailableLLMModelsResponse
	GetStatusCode() *int32
	SetBody(v *GetAvailableLLMModelsResponseBody) *GetAvailableLLMModelsResponse
	GetBody() *GetAvailableLLMModelsResponseBody
}

type GetAvailableLLMModelsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAvailableLLMModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAvailableLLMModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableLLMModelsResponse) GoString() string {
	return s.String()
}

func (s *GetAvailableLLMModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAvailableLLMModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAvailableLLMModelsResponse) GetBody() *GetAvailableLLMModelsResponseBody {
	return s.Body
}

func (s *GetAvailableLLMModelsResponse) SetHeaders(v map[string]*string) *GetAvailableLLMModelsResponse {
	s.Headers = v
	return s
}

func (s *GetAvailableLLMModelsResponse) SetStatusCode(v int32) *GetAvailableLLMModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAvailableLLMModelsResponse) SetBody(v *GetAvailableLLMModelsResponseBody) *GetAvailableLLMModelsResponse {
	s.Body = v
	return s
}

func (s *GetAvailableLLMModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
