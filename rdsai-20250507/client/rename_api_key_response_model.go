// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *RenameApiKeyResponseBody) *RenameApiKeyResponse
	GetBody() *RenameApiKeyResponseBody
}

type RenameApiKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameApiKeyResponse) GoString() string {
	return s.String()
}

func (s *RenameApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameApiKeyResponse) GetBody() *RenameApiKeyResponseBody {
	return s.Body
}

func (s *RenameApiKeyResponse) SetHeaders(v map[string]*string) *RenameApiKeyResponse {
	s.Headers = v
	return s
}

func (s *RenameApiKeyResponse) SetStatusCode(v int32) *RenameApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameApiKeyResponse) SetBody(v *RenameApiKeyResponseBody) *RenameApiKeyResponse {
	s.Body = v
	return s
}

func (s *RenameApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
