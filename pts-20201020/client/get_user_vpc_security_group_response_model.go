// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserVpcSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserVpcSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetUserVpcSecurityGroupResponseBody) *GetUserVpcSecurityGroupResponse
	GetBody() *GetUserVpcSecurityGroupResponseBody
}

type GetUserVpcSecurityGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserVpcSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserVpcSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *GetUserVpcSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserVpcSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserVpcSecurityGroupResponse) GetBody() *GetUserVpcSecurityGroupResponseBody {
	return s.Body
}

func (s *GetUserVpcSecurityGroupResponse) SetHeaders(v map[string]*string) *GetUserVpcSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *GetUserVpcSecurityGroupResponse) SetStatusCode(v int32) *GetUserVpcSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserVpcSecurityGroupResponse) SetBody(v *GetUserVpcSecurityGroupResponseBody) *GetUserVpcSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *GetUserVpcSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
