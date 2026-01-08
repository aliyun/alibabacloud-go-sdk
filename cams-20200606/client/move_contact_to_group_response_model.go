// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveContactToGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveContactToGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveContactToGroupResponse
	GetStatusCode() *int32
	SetBody(v *MoveContactToGroupResponseBody) *MoveContactToGroupResponse
	GetBody() *MoveContactToGroupResponseBody
}

type MoveContactToGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveContactToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveContactToGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveContactToGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveContactToGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveContactToGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveContactToGroupResponse) GetBody() *MoveContactToGroupResponseBody {
	return s.Body
}

func (s *MoveContactToGroupResponse) SetHeaders(v map[string]*string) *MoveContactToGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveContactToGroupResponse) SetStatusCode(v int32) *MoveContactToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveContactToGroupResponse) SetBody(v *MoveContactToGroupResponseBody) *MoveContactToGroupResponse {
	s.Body = v
	return s
}

func (s *MoveContactToGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
