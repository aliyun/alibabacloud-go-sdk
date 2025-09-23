// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourceSharePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateResourceSharePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateResourceSharePermissionResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateResourceSharePermissionResponseBody) *DisassociateResourceSharePermissionResponse
	GetBody() *DisassociateResourceSharePermissionResponseBody
}

type DisassociateResourceSharePermissionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateResourceSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateResourceSharePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourceSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *DisassociateResourceSharePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateResourceSharePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateResourceSharePermissionResponse) GetBody() *DisassociateResourceSharePermissionResponseBody {
	return s.Body
}

func (s *DisassociateResourceSharePermissionResponse) SetHeaders(v map[string]*string) *DisassociateResourceSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *DisassociateResourceSharePermissionResponse) SetStatusCode(v int32) *DisassociateResourceSharePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateResourceSharePermissionResponse) SetBody(v *DisassociateResourceSharePermissionResponseBody) *DisassociateResourceSharePermissionResponse {
	s.Body = v
	return s
}

func (s *DisassociateResourceSharePermissionResponse) Validate() error {
	return dara.Validate(s)
}
