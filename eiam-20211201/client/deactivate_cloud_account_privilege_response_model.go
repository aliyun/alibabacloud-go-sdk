// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateCloudAccountPrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeactivateCloudAccountPrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeactivateCloudAccountPrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *DeactivateCloudAccountPrivilegeResponseBody) *DeactivateCloudAccountPrivilegeResponse
	GetBody() *DeactivateCloudAccountPrivilegeResponseBody
}

type DeactivateCloudAccountPrivilegeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeactivateCloudAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeactivateCloudAccountPrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeactivateCloudAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *DeactivateCloudAccountPrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeactivateCloudAccountPrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeactivateCloudAccountPrivilegeResponse) GetBody() *DeactivateCloudAccountPrivilegeResponseBody {
	return s.Body
}

func (s *DeactivateCloudAccountPrivilegeResponse) SetHeaders(v map[string]*string) *DeactivateCloudAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *DeactivateCloudAccountPrivilegeResponse) SetStatusCode(v int32) *DeactivateCloudAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactivateCloudAccountPrivilegeResponse) SetBody(v *DeactivateCloudAccountPrivilegeResponseBody) *DeactivateCloudAccountPrivilegeResponse {
	s.Body = v
	return s
}

func (s *DeactivateCloudAccountPrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
