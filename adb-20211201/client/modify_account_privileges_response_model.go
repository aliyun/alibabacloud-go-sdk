// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountPrivilegesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountPrivilegesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountPrivilegesResponseBody) *ModifyAccountPrivilegesResponse
	GetBody() *ModifyAccountPrivilegesResponseBody
}

type ModifyAccountPrivilegesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountPrivilegesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountPrivilegesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountPrivilegesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountPrivilegesResponse) GetBody() *ModifyAccountPrivilegesResponseBody {
	return s.Body
}

func (s *ModifyAccountPrivilegesResponse) SetHeaders(v map[string]*string) *ModifyAccountPrivilegesResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPrivilegesResponse) SetStatusCode(v int32) *ModifyAccountPrivilegesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPrivilegesResponse) SetBody(v *ModifyAccountPrivilegesResponseBody) *ModifyAccountPrivilegesResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountPrivilegesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
