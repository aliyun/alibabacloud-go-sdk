// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSshKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetSshKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetSshKeyResponse
	GetStatusCode() *int32
	SetBody(v *ResetSshKeyResponseBody) *ResetSshKeyResponse
	GetBody() *ResetSshKeyResponseBody
}

type ResetSshKeyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSshKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSshKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetSshKeyResponse) GoString() string {
	return s.String()
}

func (s *ResetSshKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetSshKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetSshKeyResponse) GetBody() *ResetSshKeyResponseBody {
	return s.Body
}

func (s *ResetSshKeyResponse) SetHeaders(v map[string]*string) *ResetSshKeyResponse {
	s.Headers = v
	return s
}

func (s *ResetSshKeyResponse) SetStatusCode(v int32) *ResetSshKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSshKeyResponse) SetBody(v *ResetSshKeyResponseBody) *ResetSshKeyResponse {
	s.Body = v
	return s
}

func (s *ResetSshKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
