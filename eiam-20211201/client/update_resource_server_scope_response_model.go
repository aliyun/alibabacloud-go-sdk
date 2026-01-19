// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceServerScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceServerScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceServerScopeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceServerScopeResponseBody) *UpdateResourceServerScopeResponse
	GetBody() *UpdateResourceServerScopeResponseBody
}

type UpdateResourceServerScopeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceServerScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceServerScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceServerScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceServerScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceServerScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceServerScopeResponse) GetBody() *UpdateResourceServerScopeResponseBody {
	return s.Body
}

func (s *UpdateResourceServerScopeResponse) SetHeaders(v map[string]*string) *UpdateResourceServerScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceServerScopeResponse) SetStatusCode(v int32) *UpdateResourceServerScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceServerScopeResponse) SetBody(v *UpdateResourceServerScopeResponseBody) *UpdateResourceServerScopeResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceServerScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
