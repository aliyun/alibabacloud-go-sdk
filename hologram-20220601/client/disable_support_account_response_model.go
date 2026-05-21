// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSupportAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSupportAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSupportAccountResponse
	GetStatusCode() *int32
	SetBody(v *DisableSupportAccountResponseBody) *DisableSupportAccountResponse
	GetBody() *DisableSupportAccountResponseBody
}

type DisableSupportAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSupportAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSupportAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSupportAccountResponse) GoString() string {
	return s.String()
}

func (s *DisableSupportAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSupportAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSupportAccountResponse) GetBody() *DisableSupportAccountResponseBody {
	return s.Body
}

func (s *DisableSupportAccountResponse) SetHeaders(v map[string]*string) *DisableSupportAccountResponse {
	s.Headers = v
	return s
}

func (s *DisableSupportAccountResponse) SetStatusCode(v int32) *DisableSupportAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSupportAccountResponse) SetBody(v *DisableSupportAccountResponseBody) *DisableSupportAccountResponse {
	s.Body = v
	return s
}

func (s *DisableSupportAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
