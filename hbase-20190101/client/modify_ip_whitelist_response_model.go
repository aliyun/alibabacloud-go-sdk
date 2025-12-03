// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpWhitelistResponseBody) *ModifyIpWhitelistResponse
	GetBody() *ModifyIpWhitelistResponseBody
}

type ModifyIpWhitelistResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpWhitelistResponse) GetBody() *ModifyIpWhitelistResponseBody {
	return s.Body
}

func (s *ModifyIpWhitelistResponse) SetHeaders(v map[string]*string) *ModifyIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpWhitelistResponse) SetStatusCode(v int32) *ModifyIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpWhitelistResponse) SetBody(v *ModifyIpWhitelistResponseBody) *ModifyIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *ModifyIpWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
