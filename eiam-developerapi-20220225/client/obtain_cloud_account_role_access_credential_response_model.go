// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCloudAccountRoleAccessCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ObtainCloudAccountRoleAccessCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ObtainCloudAccountRoleAccessCredentialResponse
	GetStatusCode() *int32
	SetBody(v *ObtainCloudAccountRoleAccessCredentialResponseBody) *ObtainCloudAccountRoleAccessCredentialResponse
	GetBody() *ObtainCloudAccountRoleAccessCredentialResponseBody
}

type ObtainCloudAccountRoleAccessCredentialResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainCloudAccountRoleAccessCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainCloudAccountRoleAccessCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialResponse) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ObtainCloudAccountRoleAccessCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ObtainCloudAccountRoleAccessCredentialResponse) GetBody() *ObtainCloudAccountRoleAccessCredentialResponseBody {
	return s.Body
}

func (s *ObtainCloudAccountRoleAccessCredentialResponse) SetHeaders(v map[string]*string) *ObtainCloudAccountRoleAccessCredentialResponse {
	s.Headers = v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponse) SetStatusCode(v int32) *ObtainCloudAccountRoleAccessCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponse) SetBody(v *ObtainCloudAccountRoleAccessCredentialResponseBody) *ObtainCloudAccountRoleAccessCredentialResponse {
	s.Body = v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
