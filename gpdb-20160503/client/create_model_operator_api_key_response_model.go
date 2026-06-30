// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelOperatorApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelOperatorApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelOperatorApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelOperatorApiKeyResponseBody) *CreateModelOperatorApiKeyResponse
	GetBody() *CreateModelOperatorApiKeyResponseBody
}

type CreateModelOperatorApiKeyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelOperatorApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelOperatorApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelOperatorApiKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateModelOperatorApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelOperatorApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelOperatorApiKeyResponse) GetBody() *CreateModelOperatorApiKeyResponseBody {
	return s.Body
}

func (s *CreateModelOperatorApiKeyResponse) SetHeaders(v map[string]*string) *CreateModelOperatorApiKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateModelOperatorApiKeyResponse) SetStatusCode(v int32) *CreateModelOperatorApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelOperatorApiKeyResponse) SetBody(v *CreateModelOperatorApiKeyResponseBody) *CreateModelOperatorApiKeyResponse {
	s.Body = v
	return s
}

func (s *CreateModelOperatorApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
