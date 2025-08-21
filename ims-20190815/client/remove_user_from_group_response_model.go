// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserFromGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserFromGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserFromGroupResponseBody) *RemoveUserFromGroupResponse
	GetBody() *RemoveUserFromGroupResponseBody
}

type RemoveUserFromGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserFromGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserFromGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserFromGroupResponse) GetBody() *RemoveUserFromGroupResponseBody {
	return s.Body
}

func (s *RemoveUserFromGroupResponse) SetHeaders(v map[string]*string) *RemoveUserFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromGroupResponse) SetStatusCode(v int32) *RemoveUserFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromGroupResponse) SetBody(v *RemoveUserFromGroupResponseBody) *RemoveUserFromGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveUserFromGroupResponse) Validate() error {
	return dara.Validate(s)
}
