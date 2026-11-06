// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateCloudAccountPrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ActivateCloudAccountPrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ActivateCloudAccountPrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *ActivateCloudAccountPrivilegeResponseBody) *ActivateCloudAccountPrivilegeResponse
	GetBody() *ActivateCloudAccountPrivilegeResponseBody
}

type ActivateCloudAccountPrivilegeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateCloudAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateCloudAccountPrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s ActivateCloudAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *ActivateCloudAccountPrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ActivateCloudAccountPrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ActivateCloudAccountPrivilegeResponse) GetBody() *ActivateCloudAccountPrivilegeResponseBody {
	return s.Body
}

func (s *ActivateCloudAccountPrivilegeResponse) SetHeaders(v map[string]*string) *ActivateCloudAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *ActivateCloudAccountPrivilegeResponse) SetStatusCode(v int32) *ActivateCloudAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateCloudAccountPrivilegeResponse) SetBody(v *ActivateCloudAccountPrivilegeResponseBody) *ActivateCloudAccountPrivilegeResponse {
	s.Body = v
	return s
}

func (s *ActivateCloudAccountPrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
