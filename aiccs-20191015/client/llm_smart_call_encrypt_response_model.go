// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallEncryptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LlmSmartCallEncryptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LlmSmartCallEncryptResponse
	GetStatusCode() *int32
	SetBody(v *LlmSmartCallEncryptResponseBody) *LlmSmartCallEncryptResponse
	GetBody() *LlmSmartCallEncryptResponseBody
}

type LlmSmartCallEncryptResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LlmSmartCallEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LlmSmartCallEncryptResponse) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallEncryptResponse) GoString() string {
	return s.String()
}

func (s *LlmSmartCallEncryptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LlmSmartCallEncryptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LlmSmartCallEncryptResponse) GetBody() *LlmSmartCallEncryptResponseBody {
	return s.Body
}

func (s *LlmSmartCallEncryptResponse) SetHeaders(v map[string]*string) *LlmSmartCallEncryptResponse {
	s.Headers = v
	return s
}

func (s *LlmSmartCallEncryptResponse) SetStatusCode(v int32) *LlmSmartCallEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *LlmSmartCallEncryptResponse) SetBody(v *LlmSmartCallEncryptResponseBody) *LlmSmartCallEncryptResponse {
	s.Body = v
	return s
}

func (s *LlmSmartCallEncryptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
