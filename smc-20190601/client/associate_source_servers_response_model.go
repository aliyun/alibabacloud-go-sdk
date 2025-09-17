// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateSourceServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateSourceServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateSourceServersResponse
	GetStatusCode() *int32
	SetBody(v *AssociateSourceServersResponseBody) *AssociateSourceServersResponse
	GetBody() *AssociateSourceServersResponseBody
}

type AssociateSourceServersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateSourceServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateSourceServersResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateSourceServersResponse) GoString() string {
	return s.String()
}

func (s *AssociateSourceServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateSourceServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateSourceServersResponse) GetBody() *AssociateSourceServersResponseBody {
	return s.Body
}

func (s *AssociateSourceServersResponse) SetHeaders(v map[string]*string) *AssociateSourceServersResponse {
	s.Headers = v
	return s
}

func (s *AssociateSourceServersResponse) SetStatusCode(v int32) *AssociateSourceServersResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateSourceServersResponse) SetBody(v *AssociateSourceServersResponseBody) *AssociateSourceServersResponse {
	s.Body = v
	return s
}

func (s *AssociateSourceServersResponse) Validate() error {
	return dara.Validate(s)
}
