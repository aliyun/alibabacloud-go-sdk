// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveGroupResponseBody) *RemoveGroupResponse
	GetBody() *RemoveGroupResponseBody
}

type RemoveGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveGroupResponse) GetBody() *RemoveGroupResponseBody {
	return s.Body
}

func (s *RemoveGroupResponse) SetHeaders(v map[string]*string) *RemoveGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupResponse) SetStatusCode(v int32) *RemoveGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveGroupResponse) SetBody(v *RemoveGroupResponseBody) *RemoveGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveGroupResponse) Validate() error {
	return dara.Validate(s)
}
