// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDTCSecurityIpHostsForSQLServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDTCSecurityIpHostsForSQLServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDTCSecurityIpHostsForSQLServerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDTCSecurityIpHostsForSQLServerResponseBody) *ModifyDTCSecurityIpHostsForSQLServerResponse
	GetBody() *ModifyDTCSecurityIpHostsForSQLServerResponseBody
}

type ModifyDTCSecurityIpHostsForSQLServerResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDTCSecurityIpHostsForSQLServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDTCSecurityIpHostsForSQLServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDTCSecurityIpHostsForSQLServerResponse) GoString() string {
	return s.String()
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponse) GetBody() *ModifyDTCSecurityIpHostsForSQLServerResponseBody {
	return s.Body
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponse) SetHeaders(v map[string]*string) *ModifyDTCSecurityIpHostsForSQLServerResponse {
	s.Headers = v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponse) SetStatusCode(v int32) *ModifyDTCSecurityIpHostsForSQLServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponse) SetBody(v *ModifyDTCSecurityIpHostsForSQLServerResponseBody) *ModifyDTCSecurityIpHostsForSQLServerResponse {
	s.Body = v
	return s
}

func (s *ModifyDTCSecurityIpHostsForSQLServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
