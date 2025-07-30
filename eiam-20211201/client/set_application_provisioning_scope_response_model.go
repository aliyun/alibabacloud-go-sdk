// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApplicationProvisioningScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApplicationProvisioningScopeResponse
	GetStatusCode() *int32
	SetBody(v *SetApplicationProvisioningScopeResponseBody) *SetApplicationProvisioningScopeResponse
	GetBody() *SetApplicationProvisioningScopeResponseBody
}

type SetApplicationProvisioningScopeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationProvisioningScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationProvisioningScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningScopeResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApplicationProvisioningScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApplicationProvisioningScopeResponse) GetBody() *SetApplicationProvisioningScopeResponseBody {
	return s.Body
}

func (s *SetApplicationProvisioningScopeResponse) SetHeaders(v map[string]*string) *SetApplicationProvisioningScopeResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationProvisioningScopeResponse) SetStatusCode(v int32) *SetApplicationProvisioningScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationProvisioningScopeResponse) SetBody(v *SetApplicationProvisioningScopeResponseBody) *SetApplicationProvisioningScopeResponse {
	s.Body = v
	return s
}

func (s *SetApplicationProvisioningScopeResponse) Validate() error {
	return dara.Validate(s)
}
