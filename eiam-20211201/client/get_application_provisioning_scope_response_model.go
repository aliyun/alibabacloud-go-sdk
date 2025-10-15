// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationProvisioningScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationProvisioningScopeResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationProvisioningScopeResponseBody) *GetApplicationProvisioningScopeResponse
	GetBody() *GetApplicationProvisioningScopeResponseBody
}

type GetApplicationProvisioningScopeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationProvisioningScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationProvisioningScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationProvisioningScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationProvisioningScopeResponse) GetBody() *GetApplicationProvisioningScopeResponseBody {
	return s.Body
}

func (s *GetApplicationProvisioningScopeResponse) SetHeaders(v map[string]*string) *GetApplicationProvisioningScopeResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) SetStatusCode(v int32) *GetApplicationProvisioningScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) SetBody(v *GetApplicationProvisioningScopeResponseBody) *GetApplicationProvisioningScopeResponse {
	s.Body = v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
