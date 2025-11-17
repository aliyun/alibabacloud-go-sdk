// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLLMConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLLMConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLLMConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetLLMConfigResponseBody) *GetLLMConfigResponse
	GetBody() *GetLLMConfigResponseBody
}

type GetLLMConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLLMConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLLMConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLLMConfigResponse) GoString() string {
	return s.String()
}

func (s *GetLLMConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLLMConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLLMConfigResponse) GetBody() *GetLLMConfigResponseBody {
	return s.Body
}

func (s *GetLLMConfigResponse) SetHeaders(v map[string]*string) *GetLLMConfigResponse {
	s.Headers = v
	return s
}

func (s *GetLLMConfigResponse) SetStatusCode(v int32) *GetLLMConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLLMConfigResponse) SetBody(v *GetLLMConfigResponseBody) *GetLLMConfigResponse {
	s.Body = v
	return s
}

func (s *GetLLMConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
