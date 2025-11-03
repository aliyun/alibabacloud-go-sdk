// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceSSLResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceSSLResponseBody) *ModifyDBInstanceSSLResponse
	GetBody() *ModifyDBInstanceSSLResponseBody
}

type ModifyDBInstanceSSLResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceSSLResponse) GetBody() *ModifyDBInstanceSSLResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceSSLResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceSSLResponse) SetStatusCode(v int32) *ModifyDBInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceSSLResponse) SetBody(v *ModifyDBInstanceSSLResponseBody) *ModifyDBInstanceSSLResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceSSLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
