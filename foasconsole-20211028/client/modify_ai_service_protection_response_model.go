// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAiServiceProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAiServiceProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAiServiceProtectionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAiServiceProtectionResponseBody) *ModifyAiServiceProtectionResponse
	GetBody() *ModifyAiServiceProtectionResponseBody
}

type ModifyAiServiceProtectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAiServiceProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAiServiceProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAiServiceProtectionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAiServiceProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAiServiceProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAiServiceProtectionResponse) GetBody() *ModifyAiServiceProtectionResponseBody {
	return s.Body
}

func (s *ModifyAiServiceProtectionResponse) SetHeaders(v map[string]*string) *ModifyAiServiceProtectionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAiServiceProtectionResponse) SetStatusCode(v int32) *ModifyAiServiceProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAiServiceProtectionResponse) SetBody(v *ModifyAiServiceProtectionResponseBody) *ModifyAiServiceProtectionResponse {
	s.Body = v
	return s
}

func (s *ModifyAiServiceProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
