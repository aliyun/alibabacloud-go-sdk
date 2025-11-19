// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityIPGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityIPGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityIPGroupResponseBody) *ModifySecurityIPGroupResponse
	GetBody() *ModifySecurityIPGroupResponseBody
}

type ModifySecurityIPGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityIPGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityIPGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityIPGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityIPGroupResponse) GetBody() *ModifySecurityIPGroupResponseBody {
	return s.Body
}

func (s *ModifySecurityIPGroupResponse) SetHeaders(v map[string]*string) *ModifySecurityIPGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIPGroupResponse) SetStatusCode(v int32) *ModifySecurityIPGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIPGroupResponse) SetBody(v *ModifySecurityIPGroupResponseBody) *ModifySecurityIPGroupResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityIPGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
