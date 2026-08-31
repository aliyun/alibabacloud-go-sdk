// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmModelProviderByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLlmModelProviderByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLlmModelProviderByNameResponse
	GetStatusCode() *int32
	SetBody(v *GetLlmModelProviderByNameResponseBody) *GetLlmModelProviderByNameResponse
	GetBody() *GetLlmModelProviderByNameResponseBody
}

type GetLlmModelProviderByNameResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLlmModelProviderByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLlmModelProviderByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProviderByNameResponse) GoString() string {
	return s.String()
}

func (s *GetLlmModelProviderByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLlmModelProviderByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLlmModelProviderByNameResponse) GetBody() *GetLlmModelProviderByNameResponseBody {
	return s.Body
}

func (s *GetLlmModelProviderByNameResponse) SetHeaders(v map[string]*string) *GetLlmModelProviderByNameResponse {
	s.Headers = v
	return s
}

func (s *GetLlmModelProviderByNameResponse) SetStatusCode(v int32) *GetLlmModelProviderByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLlmModelProviderByNameResponse) SetBody(v *GetLlmModelProviderByNameResponseBody) *GetLlmModelProviderByNameResponse {
	s.Body = v
	return s
}

func (s *GetLlmModelProviderByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
