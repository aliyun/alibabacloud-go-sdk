// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApiKeyResponseBody) *DeleteApiKeyResponse
	GetBody() *DeleteApiKeyResponseBody
}

type DeleteApiKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApiKeyResponse) GetBody() *DeleteApiKeyResponseBody {
	return s.Body
}

func (s *DeleteApiKeyResponse) SetHeaders(v map[string]*string) *DeleteApiKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiKeyResponse) SetStatusCode(v int32) *DeleteApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiKeyResponse) SetBody(v *DeleteApiKeyResponseBody) *DeleteApiKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
