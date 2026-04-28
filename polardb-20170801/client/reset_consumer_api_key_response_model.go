// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetConsumerApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetConsumerApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetConsumerApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ResetConsumerApiKeyResponseBody) *ResetConsumerApiKeyResponse
	GetBody() *ResetConsumerApiKeyResponseBody
}

type ResetConsumerApiKeyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetConsumerApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetConsumerApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetConsumerApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ResetConsumerApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetConsumerApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetConsumerApiKeyResponse) GetBody() *ResetConsumerApiKeyResponseBody {
	return s.Body
}

func (s *ResetConsumerApiKeyResponse) SetHeaders(v map[string]*string) *ResetConsumerApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ResetConsumerApiKeyResponse) SetStatusCode(v int32) *ResetConsumerApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetConsumerApiKeyResponse) SetBody(v *ResetConsumerApiKeyResponseBody) *ResetConsumerApiKeyResponse {
	s.Body = v
	return s
}

func (s *ResetConsumerApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
