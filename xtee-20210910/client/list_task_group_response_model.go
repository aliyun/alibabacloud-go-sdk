// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskGroupResponseBody) *ListTaskGroupResponse
	GetBody() *ListTaskGroupResponseBody
}

type ListTaskGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *ListTaskGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskGroupResponse) GetBody() *ListTaskGroupResponseBody {
	return s.Body
}

func (s *ListTaskGroupResponse) SetHeaders(v map[string]*string) *ListTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *ListTaskGroupResponse) SetStatusCode(v int32) *ListTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskGroupResponse) SetBody(v *ListTaskGroupResponseBody) *ListTaskGroupResponse {
	s.Body = v
	return s
}

func (s *ListTaskGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
