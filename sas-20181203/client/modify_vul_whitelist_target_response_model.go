// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulWhitelistTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVulWhitelistTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVulWhitelistTargetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVulWhitelistTargetResponseBody) *ModifyVulWhitelistTargetResponse
	GetBody() *ModifyVulWhitelistTargetResponseBody
}

type ModifyVulWhitelistTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVulWhitelistTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVulWhitelistTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulWhitelistTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyVulWhitelistTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVulWhitelistTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVulWhitelistTargetResponse) GetBody() *ModifyVulWhitelistTargetResponseBody {
	return s.Body
}

func (s *ModifyVulWhitelistTargetResponse) SetHeaders(v map[string]*string) *ModifyVulWhitelistTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyVulWhitelistTargetResponse) SetStatusCode(v int32) *ModifyVulWhitelistTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVulWhitelistTargetResponse) SetBody(v *ModifyVulWhitelistTargetResponseBody) *ModifyVulWhitelistTargetResponse {
	s.Body = v
	return s
}

func (s *ModifyVulWhitelistTargetResponse) Validate() error {
	return dara.Validate(s)
}
