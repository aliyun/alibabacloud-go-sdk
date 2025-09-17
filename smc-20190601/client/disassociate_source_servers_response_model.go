// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateSourceServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateSourceServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateSourceServersResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateSourceServersResponseBody) *DisassociateSourceServersResponse
	GetBody() *DisassociateSourceServersResponseBody
}

type DisassociateSourceServersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateSourceServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateSourceServersResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateSourceServersResponse) GoString() string {
	return s.String()
}

func (s *DisassociateSourceServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateSourceServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateSourceServersResponse) GetBody() *DisassociateSourceServersResponseBody {
	return s.Body
}

func (s *DisassociateSourceServersResponse) SetHeaders(v map[string]*string) *DisassociateSourceServersResponse {
	s.Headers = v
	return s
}

func (s *DisassociateSourceServersResponse) SetStatusCode(v int32) *DisassociateSourceServersResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateSourceServersResponse) SetBody(v *DisassociateSourceServersResponseBody) *DisassociateSourceServersResponse {
	s.Body = v
	return s
}

func (s *DisassociateSourceServersResponse) Validate() error {
	return dara.Validate(s)
}
