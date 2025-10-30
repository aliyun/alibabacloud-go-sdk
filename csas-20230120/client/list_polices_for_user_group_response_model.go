// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicesForUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicesForUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicesForUserGroupResponseBody) *ListPolicesForUserGroupResponse
	GetBody() *ListPolicesForUserGroupResponseBody
}

type ListPolicesForUserGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicesForUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicesForUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicesForUserGroupResponse) GetBody() *ListPolicesForUserGroupResponseBody {
	return s.Body
}

func (s *ListPolicesForUserGroupResponse) SetHeaders(v map[string]*string) *ListPolicesForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListPolicesForUserGroupResponse) SetStatusCode(v int32) *ListPolicesForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicesForUserGroupResponse) SetBody(v *ListPolicesForUserGroupResponseBody) *ListPolicesForUserGroupResponse {
	s.Body = v
	return s
}

func (s *ListPolicesForUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
