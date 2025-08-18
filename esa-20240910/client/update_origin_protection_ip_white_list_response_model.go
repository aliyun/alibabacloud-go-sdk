// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginProtectionIpWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOriginProtectionIpWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOriginProtectionIpWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOriginProtectionIpWhiteListResponseBody) *UpdateOriginProtectionIpWhiteListResponse
	GetBody() *UpdateOriginProtectionIpWhiteListResponseBody
}

type UpdateOriginProtectionIpWhiteListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOriginProtectionIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOriginProtectionIpWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginProtectionIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateOriginProtectionIpWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOriginProtectionIpWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOriginProtectionIpWhiteListResponse) GetBody() *UpdateOriginProtectionIpWhiteListResponseBody {
	return s.Body
}

func (s *UpdateOriginProtectionIpWhiteListResponse) SetHeaders(v map[string]*string) *UpdateOriginProtectionIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateOriginProtectionIpWhiteListResponse) SetStatusCode(v int32) *UpdateOriginProtectionIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOriginProtectionIpWhiteListResponse) SetBody(v *UpdateOriginProtectionIpWhiteListResponseBody) *UpdateOriginProtectionIpWhiteListResponse {
	s.Body = v
	return s
}

func (s *UpdateOriginProtectionIpWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
