// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateAclsFromListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateAclsFromListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateAclsFromListenerResponse
	GetStatusCode() *int32
	SetBody(v *DissociateAclsFromListenerResponseBody) *DissociateAclsFromListenerResponse
	GetBody() *DissociateAclsFromListenerResponseBody
}

type DissociateAclsFromListenerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateAclsFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateAclsFromListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateAclsFromListenerResponse) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateAclsFromListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateAclsFromListenerResponse) GetBody() *DissociateAclsFromListenerResponseBody {
	return s.Body
}

func (s *DissociateAclsFromListenerResponse) SetHeaders(v map[string]*string) *DissociateAclsFromListenerResponse {
	s.Headers = v
	return s
}

func (s *DissociateAclsFromListenerResponse) SetStatusCode(v int32) *DissociateAclsFromListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateAclsFromListenerResponse) SetBody(v *DissociateAclsFromListenerResponseBody) *DissociateAclsFromListenerResponse {
	s.Body = v
	return s
}

func (s *DissociateAclsFromListenerResponse) Validate() error {
	return dara.Validate(s)
}
