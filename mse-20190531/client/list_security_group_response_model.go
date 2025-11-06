// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListSecurityGroupResponseBody) *ListSecurityGroupResponse
	GetBody() *ListSecurityGroupResponseBody
}

type ListSecurityGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecurityGroupResponse) GetBody() *ListSecurityGroupResponseBody {
	return s.Body
}

func (s *ListSecurityGroupResponse) SetHeaders(v map[string]*string) *ListSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityGroupResponse) SetStatusCode(v int32) *ListSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityGroupResponse) SetBody(v *ListSecurityGroupResponseBody) *ListSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *ListSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
