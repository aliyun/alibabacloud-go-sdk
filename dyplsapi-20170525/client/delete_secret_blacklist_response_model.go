// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretBlacklistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecretBlacklistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecretBlacklistResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecretBlacklistResponseBody) *DeleteSecretBlacklistResponse
	GetBody() *DeleteSecretBlacklistResponseBody
}

type DeleteSecretBlacklistResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretBlacklistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretBlacklistResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretBlacklistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecretBlacklistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecretBlacklistResponse) GetBody() *DeleteSecretBlacklistResponseBody {
	return s.Body
}

func (s *DeleteSecretBlacklistResponse) SetHeaders(v map[string]*string) *DeleteSecretBlacklistResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretBlacklistResponse) SetStatusCode(v int32) *DeleteSecretBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretBlacklistResponse) SetBody(v *DeleteSecretBlacklistResponseBody) *DeleteSecretBlacklistResponse {
	s.Body = v
	return s
}

func (s *DeleteSecretBlacklistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
