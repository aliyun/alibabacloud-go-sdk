// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterSecurityIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterSecurityIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterSecurityIpListResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterSecurityIpListResponseBody) *ModifyClusterSecurityIpListResponse
	GetBody() *ModifyClusterSecurityIpListResponseBody
}

type ModifyClusterSecurityIpListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterSecurityIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterSecurityIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterSecurityIpListResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterSecurityIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterSecurityIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterSecurityIpListResponse) GetBody() *ModifyClusterSecurityIpListResponseBody {
	return s.Body
}

func (s *ModifyClusterSecurityIpListResponse) SetHeaders(v map[string]*string) *ModifyClusterSecurityIpListResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterSecurityIpListResponse) SetStatusCode(v int32) *ModifyClusterSecurityIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterSecurityIpListResponse) SetBody(v *ModifyClusterSecurityIpListResponseBody) *ModifyClusterSecurityIpListResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterSecurityIpListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
