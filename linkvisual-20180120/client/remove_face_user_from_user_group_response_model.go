// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFaceUserFromUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveFaceUserFromUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveFaceUserFromUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveFaceUserFromUserGroupResponseBody) *RemoveFaceUserFromUserGroupResponse
	GetBody() *RemoveFaceUserFromUserGroupResponseBody
}

type RemoveFaceUserFromUserGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFaceUserFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFaceUserFromUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveFaceUserFromUserGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveFaceUserFromUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveFaceUserFromUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveFaceUserFromUserGroupResponse) GetBody() *RemoveFaceUserFromUserGroupResponseBody {
	return s.Body
}

func (s *RemoveFaceUserFromUserGroupResponse) SetHeaders(v map[string]*string) *RemoveFaceUserFromUserGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponse) SetStatusCode(v int32) *RemoveFaceUserFromUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponse) SetBody(v *RemoveFaceUserFromUserGroupResponseBody) *RemoveFaceUserFromUserGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveFaceUserFromUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
