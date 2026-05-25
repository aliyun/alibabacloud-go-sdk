// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallFullDuplexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LlmSmartCallFullDuplexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LlmSmartCallFullDuplexResponse
	GetStatusCode() *int32
	SetBody(v *LlmSmartCallFullDuplexResponseBody) *LlmSmartCallFullDuplexResponse
	GetBody() *LlmSmartCallFullDuplexResponseBody
}

type LlmSmartCallFullDuplexResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LlmSmartCallFullDuplexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LlmSmartCallFullDuplexResponse) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallFullDuplexResponse) GoString() string {
	return s.String()
}

func (s *LlmSmartCallFullDuplexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LlmSmartCallFullDuplexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LlmSmartCallFullDuplexResponse) GetBody() *LlmSmartCallFullDuplexResponseBody {
	return s.Body
}

func (s *LlmSmartCallFullDuplexResponse) SetHeaders(v map[string]*string) *LlmSmartCallFullDuplexResponse {
	s.Headers = v
	return s
}

func (s *LlmSmartCallFullDuplexResponse) SetStatusCode(v int32) *LlmSmartCallFullDuplexResponse {
	s.StatusCode = &v
	return s
}

func (s *LlmSmartCallFullDuplexResponse) SetBody(v *LlmSmartCallFullDuplexResponseBody) *LlmSmartCallFullDuplexResponse {
	s.Body = v
	return s
}

func (s *LlmSmartCallFullDuplexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
