// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourceShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateResourceShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateResourceShareResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateResourceShareResponseBody) *DisassociateResourceShareResponse
	GetBody() *DisassociateResourceShareResponseBody
}

type DisassociateResourceShareResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateResourceShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateResourceShareResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourceShareResponse) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateResourceShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateResourceShareResponse) GetBody() *DisassociateResourceShareResponseBody {
	return s.Body
}

func (s *DisassociateResourceShareResponse) SetHeaders(v map[string]*string) *DisassociateResourceShareResponse {
	s.Headers = v
	return s
}

func (s *DisassociateResourceShareResponse) SetStatusCode(v int32) *DisassociateResourceShareResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateResourceShareResponse) SetBody(v *DisassociateResourceShareResponseBody) *DisassociateResourceShareResponse {
	s.Body = v
	return s
}

func (s *DisassociateResourceShareResponse) Validate() error {
	return dara.Validate(s)
}
