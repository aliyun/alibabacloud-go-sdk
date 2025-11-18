// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceIpWhitelistResponseBody) *ModifyInstanceIpWhitelistResponse
	GetBody() *ModifyInstanceIpWhitelistResponseBody
}

type ModifyInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceIpWhitelistResponse) GetBody() *ModifyInstanceIpWhitelistResponseBody {
	return s.Body
}

func (s *ModifyInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *ModifyInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceIpWhitelistResponse) SetStatusCode(v int32) *ModifyInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceIpWhitelistResponse) SetBody(v *ModifyInstanceIpWhitelistResponseBody) *ModifyInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceIpWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
