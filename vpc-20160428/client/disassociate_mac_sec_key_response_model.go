// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateMacSecKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateMacSecKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateMacSecKeyResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateMacSecKeyResponseBody) *DisassociateMacSecKeyResponse
	GetBody() *DisassociateMacSecKeyResponseBody
}

type DisassociateMacSecKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateMacSecKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateMacSecKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateMacSecKeyResponse) GoString() string {
	return s.String()
}

func (s *DisassociateMacSecKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateMacSecKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateMacSecKeyResponse) GetBody() *DisassociateMacSecKeyResponseBody {
	return s.Body
}

func (s *DisassociateMacSecKeyResponse) SetHeaders(v map[string]*string) *DisassociateMacSecKeyResponse {
	s.Headers = v
	return s
}

func (s *DisassociateMacSecKeyResponse) SetStatusCode(v int32) *DisassociateMacSecKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateMacSecKeyResponse) SetBody(v *DisassociateMacSecKeyResponseBody) *DisassociateMacSecKeyResponse {
	s.Body = v
	return s
}

func (s *DisassociateMacSecKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
