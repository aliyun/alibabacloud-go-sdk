// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMessageGroupResponseBody) *UpdateMessageGroupResponse
	GetBody() *UpdateMessageGroupResponseBody
}

type UpdateMessageGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMessageGroupResponse) GetBody() *UpdateMessageGroupResponseBody {
	return s.Body
}

func (s *UpdateMessageGroupResponse) SetHeaders(v map[string]*string) *UpdateMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageGroupResponse) SetStatusCode(v int32) *UpdateMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageGroupResponse) SetBody(v *UpdateMessageGroupResponseBody) *UpdateMessageGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateMessageGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
