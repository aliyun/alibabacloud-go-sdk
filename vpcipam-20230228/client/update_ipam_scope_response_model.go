// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpamScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpamScopeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpamScopeResponseBody) *UpdateIpamScopeResponse
	GetBody() *UpdateIpamScopeResponseBody
}

type UpdateIpamScopeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpamScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpamScopeResponse) GetBody() *UpdateIpamScopeResponseBody {
	return s.Body
}

func (s *UpdateIpamScopeResponse) SetHeaders(v map[string]*string) *UpdateIpamScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamScopeResponse) SetStatusCode(v int32) *UpdateIpamScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamScopeResponse) SetBody(v *UpdateIpamScopeResponseBody) *UpdateIpamScopeResponse {
	s.Body = v
	return s
}

func (s *UpdateIpamScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
