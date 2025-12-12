// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAssignedSessionGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevertAssignedSessionGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevertAssignedSessionGroupResponse
	GetStatusCode() *int32
	SetBody(v *RevertAssignedSessionGroupResponseBody) *RevertAssignedSessionGroupResponse
	GetBody() *RevertAssignedSessionGroupResponseBody
}

type RevertAssignedSessionGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertAssignedSessionGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertAssignedSessionGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RevertAssignedSessionGroupResponse) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevertAssignedSessionGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevertAssignedSessionGroupResponse) GetBody() *RevertAssignedSessionGroupResponseBody {
	return s.Body
}

func (s *RevertAssignedSessionGroupResponse) SetHeaders(v map[string]*string) *RevertAssignedSessionGroupResponse {
	s.Headers = v
	return s
}

func (s *RevertAssignedSessionGroupResponse) SetStatusCode(v int32) *RevertAssignedSessionGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertAssignedSessionGroupResponse) SetBody(v *RevertAssignedSessionGroupResponseBody) *RevertAssignedSessionGroupResponse {
	s.Body = v
	return s
}

func (s *RevertAssignedSessionGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
