// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSshKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSshKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSshKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateSshKeyResponseBody) *CreateSshKeyResponse
	GetBody() *CreateSshKeyResponseBody
}

type CreateSshKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSshKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSshKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSshKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateSshKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSshKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSshKeyResponse) GetBody() *CreateSshKeyResponseBody {
	return s.Body
}

func (s *CreateSshKeyResponse) SetHeaders(v map[string]*string) *CreateSshKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateSshKeyResponse) SetStatusCode(v int32) *CreateSshKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSshKeyResponse) SetBody(v *CreateSshKeyResponseBody) *CreateSshKeyResponse {
	s.Body = v
	return s
}

func (s *CreateSshKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
