// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSecretBlacklistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSecretBlacklistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSecretBlacklistResponse
	GetStatusCode() *int32
	SetBody(v *AddSecretBlacklistResponseBody) *AddSecretBlacklistResponse
	GetBody() *AddSecretBlacklistResponseBody
}

type AddSecretBlacklistResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSecretBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSecretBlacklistResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSecretBlacklistResponse) GoString() string {
	return s.String()
}

func (s *AddSecretBlacklistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSecretBlacklistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSecretBlacklistResponse) GetBody() *AddSecretBlacklistResponseBody {
	return s.Body
}

func (s *AddSecretBlacklistResponse) SetHeaders(v map[string]*string) *AddSecretBlacklistResponse {
	s.Headers = v
	return s
}

func (s *AddSecretBlacklistResponse) SetStatusCode(v int32) *AddSecretBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSecretBlacklistResponse) SetBody(v *AddSecretBlacklistResponseBody) *AddSecretBlacklistResponse {
	s.Body = v
	return s
}

func (s *AddSecretBlacklistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
