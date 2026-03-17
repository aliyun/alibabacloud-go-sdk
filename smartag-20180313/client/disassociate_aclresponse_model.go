// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateACLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateACLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateACLResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateACLResponseBody) *DisassociateACLResponse
	GetBody() *DisassociateACLResponseBody
}

type DisassociateACLResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateACLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateACLResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateACLResponse) GoString() string {
	return s.String()
}

func (s *DisassociateACLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateACLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateACLResponse) GetBody() *DisassociateACLResponseBody {
	return s.Body
}

func (s *DisassociateACLResponse) SetHeaders(v map[string]*string) *DisassociateACLResponse {
	s.Headers = v
	return s
}

func (s *DisassociateACLResponse) SetStatusCode(v int32) *DisassociateACLResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateACLResponse) SetBody(v *DisassociateACLResponseBody) *DisassociateACLResponse {
	s.Body = v
	return s
}

func (s *DisassociateACLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
