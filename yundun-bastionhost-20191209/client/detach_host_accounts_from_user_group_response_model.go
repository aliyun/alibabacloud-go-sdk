// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachHostAccountsFromUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachHostAccountsFromUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *DetachHostAccountsFromUserGroupResponseBody) *DetachHostAccountsFromUserGroupResponse
	GetBody() *DetachHostAccountsFromUserGroupResponseBody
}

type DetachHostAccountsFromUserGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachHostAccountsFromUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachHostAccountsFromUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachHostAccountsFromUserGroupResponse) GetBody() *DetachHostAccountsFromUserGroupResponseBody {
	return s.Body
}

func (s *DetachHostAccountsFromUserGroupResponse) SetHeaders(v map[string]*string) *DetachHostAccountsFromUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponse) SetStatusCode(v int32) *DetachHostAccountsFromUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponse) SetBody(v *DetachHostAccountsFromUserGroupResponseBody) *DetachHostAccountsFromUserGroupResponse {
	s.Body = v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
