// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfuseSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPxfuseSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPxfuseSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPxfuseSecurityIpsResponseBody) *ModifyPxfuseSecurityIpsResponse
	GetBody() *ModifyPxfuseSecurityIpsResponseBody
}

type ModifyPxfuseSecurityIpsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPxfuseSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPxfuseSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfuseSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyPxfuseSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPxfuseSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPxfuseSecurityIpsResponse) GetBody() *ModifyPxfuseSecurityIpsResponseBody {
	return s.Body
}

func (s *ModifyPxfuseSecurityIpsResponse) SetHeaders(v map[string]*string) *ModifyPxfuseSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponse) SetStatusCode(v int32) *ModifyPxfuseSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponse) SetBody(v *ModifyPxfuseSecurityIpsResponseBody) *ModifyPxfuseSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *ModifyPxfuseSecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
