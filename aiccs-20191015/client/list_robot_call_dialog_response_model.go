// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotCallDialogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRobotCallDialogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRobotCallDialogResponse
	GetStatusCode() *int32
	SetBody(v *ListRobotCallDialogResponseBody) *ListRobotCallDialogResponse
	GetBody() *ListRobotCallDialogResponseBody
}

type ListRobotCallDialogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRobotCallDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRobotCallDialogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRobotCallDialogResponse) GoString() string {
	return s.String()
}

func (s *ListRobotCallDialogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRobotCallDialogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRobotCallDialogResponse) GetBody() *ListRobotCallDialogResponseBody {
	return s.Body
}

func (s *ListRobotCallDialogResponse) SetHeaders(v map[string]*string) *ListRobotCallDialogResponse {
	s.Headers = v
	return s
}

func (s *ListRobotCallDialogResponse) SetStatusCode(v int32) *ListRobotCallDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRobotCallDialogResponse) SetBody(v *ListRobotCallDialogResponseBody) *ListRobotCallDialogResponse {
	s.Body = v
	return s
}

func (s *ListRobotCallDialogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
