// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCSecurityGroupPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCSecurityGroupPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCSecurityGroupPermissionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCSecurityGroupPermissionResponseBody) *ModifyRCSecurityGroupPermissionResponse
	GetBody() *ModifyRCSecurityGroupPermissionResponseBody
}

type ModifyRCSecurityGroupPermissionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCSecurityGroupPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCSecurityGroupPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCSecurityGroupPermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCSecurityGroupPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCSecurityGroupPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCSecurityGroupPermissionResponse) GetBody() *ModifyRCSecurityGroupPermissionResponseBody {
	return s.Body
}

func (s *ModifyRCSecurityGroupPermissionResponse) SetHeaders(v map[string]*string) *ModifyRCSecurityGroupPermissionResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCSecurityGroupPermissionResponse) SetStatusCode(v int32) *ModifyRCSecurityGroupPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCSecurityGroupPermissionResponse) SetBody(v *ModifyRCSecurityGroupPermissionResponseBody) *ModifyRCSecurityGroupPermissionResponse {
	s.Body = v
	return s
}

func (s *ModifyRCSecurityGroupPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
