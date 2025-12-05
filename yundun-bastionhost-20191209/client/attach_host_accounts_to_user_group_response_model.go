// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachHostAccountsToUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachHostAccountsToUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *AttachHostAccountsToUserGroupResponseBody) *AttachHostAccountsToUserGroupResponse
	GetBody() *AttachHostAccountsToUserGroupResponseBody
}

type AttachHostAccountsToUserGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachHostAccountsToUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachHostAccountsToUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachHostAccountsToUserGroupResponse) GetBody() *AttachHostAccountsToUserGroupResponseBody {
	return s.Body
}

func (s *AttachHostAccountsToUserGroupResponse) SetHeaders(v map[string]*string) *AttachHostAccountsToUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachHostAccountsToUserGroupResponse) SetStatusCode(v int32) *AttachHostAccountsToUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponse) SetBody(v *AttachHostAccountsToUserGroupResponseBody) *AttachHostAccountsToUserGroupResponse {
	s.Body = v
	return s
}

func (s *AttachHostAccountsToUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
