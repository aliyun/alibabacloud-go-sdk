// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCreateVulWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCreateVulWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCreateVulWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCreateVulWhitelistResponseBody) *ModifyCreateVulWhitelistResponse
	GetBody() *ModifyCreateVulWhitelistResponseBody
}

type ModifyCreateVulWhitelistResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCreateVulWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCreateVulWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCreateVulWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyCreateVulWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCreateVulWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCreateVulWhitelistResponse) GetBody() *ModifyCreateVulWhitelistResponseBody {
	return s.Body
}

func (s *ModifyCreateVulWhitelistResponse) SetHeaders(v map[string]*string) *ModifyCreateVulWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyCreateVulWhitelistResponse) SetStatusCode(v int32) *ModifyCreateVulWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCreateVulWhitelistResponse) SetBody(v *ModifyCreateVulWhitelistResponseBody) *ModifyCreateVulWhitelistResponse {
	s.Body = v
	return s
}

func (s *ModifyCreateVulWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
