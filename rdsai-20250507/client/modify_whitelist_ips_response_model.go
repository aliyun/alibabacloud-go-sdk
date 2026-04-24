// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhitelistIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWhitelistIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWhitelistIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWhitelistIpsResponseBody) *ModifyWhitelistIpsResponse
	GetBody() *ModifyWhitelistIpsResponseBody
}

type ModifyWhitelistIpsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWhitelistIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWhitelistIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhitelistIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifyWhitelistIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWhitelistIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWhitelistIpsResponse) GetBody() *ModifyWhitelistIpsResponseBody {
	return s.Body
}

func (s *ModifyWhitelistIpsResponse) SetHeaders(v map[string]*string) *ModifyWhitelistIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifyWhitelistIpsResponse) SetStatusCode(v int32) *ModifyWhitelistIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWhitelistIpsResponse) SetBody(v *ModifyWhitelistIpsResponseBody) *ModifyWhitelistIpsResponse {
	s.Body = v
	return s
}

func (s *ModifyWhitelistIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
