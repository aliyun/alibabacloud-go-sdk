// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMajorProtectionBlackIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMajorProtectionBlackIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMajorProtectionBlackIpResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMajorProtectionBlackIpResponseBody) *ModifyMajorProtectionBlackIpResponse
	GetBody() *ModifyMajorProtectionBlackIpResponseBody
}

type ModifyMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMajorProtectionBlackIpResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMajorProtectionBlackIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMajorProtectionBlackIpResponse) GetBody() *ModifyMajorProtectionBlackIpResponseBody {
	return s.Body
}

func (s *ModifyMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *ModifyMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *ModifyMajorProtectionBlackIpResponse) SetStatusCode(v int32) *ModifyMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpResponse) SetBody(v *ModifyMajorProtectionBlackIpResponseBody) *ModifyMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

func (s *ModifyMajorProtectionBlackIpResponse) Validate() error {
	return dara.Validate(s)
}
