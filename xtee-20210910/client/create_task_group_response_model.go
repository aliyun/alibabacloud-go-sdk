// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTaskGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTaskGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateTaskGroupResponseBody) *CreateTaskGroupResponse
	GetBody() *CreateTaskGroupResponseBody
}

type CreateTaskGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTaskGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTaskGroupResponse) GetBody() *CreateTaskGroupResponseBody {
	return s.Body
}

func (s *CreateTaskGroupResponse) SetHeaders(v map[string]*string) *CreateTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskGroupResponse) SetStatusCode(v int32) *CreateTaskGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskGroupResponse) SetBody(v *CreateTaskGroupResponseBody) *CreateTaskGroupResponse {
	s.Body = v
	return s
}

func (s *CreateTaskGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
