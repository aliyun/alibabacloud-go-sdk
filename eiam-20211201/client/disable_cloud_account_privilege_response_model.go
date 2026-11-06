// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCloudAccountPrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCloudAccountPrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCloudAccountPrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *DisableCloudAccountPrivilegeResponseBody) *DisableCloudAccountPrivilegeResponse
	GetBody() *DisableCloudAccountPrivilegeResponseBody
}

type DisableCloudAccountPrivilegeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCloudAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCloudAccountPrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCloudAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *DisableCloudAccountPrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCloudAccountPrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCloudAccountPrivilegeResponse) GetBody() *DisableCloudAccountPrivilegeResponseBody {
	return s.Body
}

func (s *DisableCloudAccountPrivilegeResponse) SetHeaders(v map[string]*string) *DisableCloudAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *DisableCloudAccountPrivilegeResponse) SetStatusCode(v int32) *DisableCloudAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCloudAccountPrivilegeResponse) SetBody(v *DisableCloudAccountPrivilegeResponseBody) *DisableCloudAccountPrivilegeResponse {
	s.Body = v
	return s
}

func (s *DisableCloudAccountPrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
