// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LlmSmartCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LlmSmartCallResponse
	GetStatusCode() *int32
	SetBody(v *LlmSmartCallResponseBody) *LlmSmartCallResponse
	GetBody() *LlmSmartCallResponseBody
}

type LlmSmartCallResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LlmSmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LlmSmartCallResponse) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallResponse) GoString() string {
	return s.String()
}

func (s *LlmSmartCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LlmSmartCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LlmSmartCallResponse) GetBody() *LlmSmartCallResponseBody {
	return s.Body
}

func (s *LlmSmartCallResponse) SetHeaders(v map[string]*string) *LlmSmartCallResponse {
	s.Headers = v
	return s
}

func (s *LlmSmartCallResponse) SetStatusCode(v int32) *LlmSmartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *LlmSmartCallResponse) SetBody(v *LlmSmartCallResponseBody) *LlmSmartCallResponse {
	s.Body = v
	return s
}

func (s *LlmSmartCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
