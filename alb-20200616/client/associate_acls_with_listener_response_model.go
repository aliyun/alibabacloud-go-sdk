// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAclsWithListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateAclsWithListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateAclsWithListenerResponse
	GetStatusCode() *int32
	SetBody(v *AssociateAclsWithListenerResponseBody) *AssociateAclsWithListenerResponse
	GetBody() *AssociateAclsWithListenerResponseBody
}

type AssociateAclsWithListenerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateAclsWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateAclsWithListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateAclsWithListenerResponse) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateAclsWithListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateAclsWithListenerResponse) GetBody() *AssociateAclsWithListenerResponseBody {
	return s.Body
}

func (s *AssociateAclsWithListenerResponse) SetHeaders(v map[string]*string) *AssociateAclsWithListenerResponse {
	s.Headers = v
	return s
}

func (s *AssociateAclsWithListenerResponse) SetStatusCode(v int32) *AssociateAclsWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateAclsWithListenerResponse) SetBody(v *AssociateAclsWithListenerResponseBody) *AssociateAclsWithListenerResponse {
	s.Body = v
	return s
}

func (s *AssociateAclsWithListenerResponse) Validate() error {
	return dara.Validate(s)
}
