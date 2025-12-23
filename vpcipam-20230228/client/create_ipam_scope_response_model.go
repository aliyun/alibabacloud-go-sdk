// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIpamScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIpamScopeResponse
	GetStatusCode() *int32
	SetBody(v *CreateIpamScopeResponseBody) *CreateIpamScopeResponse
	GetBody() *CreateIpamScopeResponseBody
}

type CreateIpamScopeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIpamScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIpamScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamScopeResponse) GoString() string {
	return s.String()
}

func (s *CreateIpamScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIpamScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIpamScopeResponse) GetBody() *CreateIpamScopeResponseBody {
	return s.Body
}

func (s *CreateIpamScopeResponse) SetHeaders(v map[string]*string) *CreateIpamScopeResponse {
	s.Headers = v
	return s
}

func (s *CreateIpamScopeResponse) SetStatusCode(v int32) *CreateIpamScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpamScopeResponse) SetBody(v *CreateIpamScopeResponseBody) *CreateIpamScopeResponse {
	s.Body = v
	return s
}

func (s *CreateIpamScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
