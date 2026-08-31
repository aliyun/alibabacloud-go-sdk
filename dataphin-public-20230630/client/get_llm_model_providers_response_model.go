// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmModelProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLlmModelProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLlmModelProvidersResponse
	GetStatusCode() *int32
	SetBody(v *GetLlmModelProvidersResponseBody) *GetLlmModelProvidersResponse
	GetBody() *GetLlmModelProvidersResponseBody
}

type GetLlmModelProvidersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLlmModelProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLlmModelProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProvidersResponse) GoString() string {
	return s.String()
}

func (s *GetLlmModelProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLlmModelProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLlmModelProvidersResponse) GetBody() *GetLlmModelProvidersResponseBody {
	return s.Body
}

func (s *GetLlmModelProvidersResponse) SetHeaders(v map[string]*string) *GetLlmModelProvidersResponse {
	s.Headers = v
	return s
}

func (s *GetLlmModelProvidersResponse) SetStatusCode(v int32) *GetLlmModelProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLlmModelProvidersResponse) SetBody(v *GetLlmModelProvidersResponseBody) *GetLlmModelProvidersResponse {
	s.Body = v
	return s
}

func (s *GetLlmModelProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
