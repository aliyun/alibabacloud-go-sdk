// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceIpWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceIpWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceIpWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceIpWhiteListResponseBody) *UpdateInstanceIpWhiteListResponse
	GetBody() *UpdateInstanceIpWhiteListResponseBody
}

type UpdateInstanceIpWhiteListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceIpWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceIpWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceIpWhiteListResponse) GetBody() *UpdateInstanceIpWhiteListResponseBody {
	return s.Body
}

func (s *UpdateInstanceIpWhiteListResponse) SetHeaders(v map[string]*string) *UpdateInstanceIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceIpWhiteListResponse) SetStatusCode(v int32) *UpdateInstanceIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceIpWhiteListResponse) SetBody(v *UpdateInstanceIpWhiteListResponseBody) *UpdateInstanceIpWhiteListResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceIpWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
