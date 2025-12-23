// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateDefaultFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateDefaultFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateDefaultFilterResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateDefaultFilterResponseBody) *DisassociateDefaultFilterResponse
	GetBody() *DisassociateDefaultFilterResponseBody
}

type DisassociateDefaultFilterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateDefaultFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateDefaultFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateDefaultFilterResponse) GoString() string {
	return s.String()
}

func (s *DisassociateDefaultFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateDefaultFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateDefaultFilterResponse) GetBody() *DisassociateDefaultFilterResponseBody {
	return s.Body
}

func (s *DisassociateDefaultFilterResponse) SetHeaders(v map[string]*string) *DisassociateDefaultFilterResponse {
	s.Headers = v
	return s
}

func (s *DisassociateDefaultFilterResponse) SetStatusCode(v int32) *DisassociateDefaultFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateDefaultFilterResponse) SetBody(v *DisassociateDefaultFilterResponseBody) *DisassociateDefaultFilterResponse {
	s.Body = v
	return s
}

func (s *DisassociateDefaultFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
