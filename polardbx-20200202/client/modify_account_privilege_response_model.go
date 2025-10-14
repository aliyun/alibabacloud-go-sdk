// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountPrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountPrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountPrivilegeResponseBody) *ModifyAccountPrivilegeResponse
	GetBody() *ModifyAccountPrivilegeResponseBody
}

type ModifyAccountPrivilegeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountPrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountPrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountPrivilegeResponse) GetBody() *ModifyAccountPrivilegeResponseBody {
	return s.Body
}

func (s *ModifyAccountPrivilegeResponse) SetHeaders(v map[string]*string) *ModifyAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPrivilegeResponse) SetStatusCode(v int32) *ModifyAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountPrivilegeResponse) SetBody(v *ModifyAccountPrivilegeResponseBody) *ModifyAccountPrivilegeResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountPrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
