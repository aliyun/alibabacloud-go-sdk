// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse
	GetBody() *DeleteAccountResponseBody
}

type DeleteAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccountResponse) GetBody() *DeleteAccountResponseBody {
	return s.Body
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetStatusCode(v int32) *DeleteAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteAccountResponse) Validate() error {
	return dara.Validate(s)
}
