// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsByGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobsByGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobsByGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListJobsByGroupResponseBody) *ListJobsByGroupResponse
	GetBody() *ListJobsByGroupResponseBody
}

type ListJobsByGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobsByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobsByGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobsByGroupResponse) GoString() string {
	return s.String()
}

func (s *ListJobsByGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobsByGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobsByGroupResponse) GetBody() *ListJobsByGroupResponseBody {
	return s.Body
}

func (s *ListJobsByGroupResponse) SetHeaders(v map[string]*string) *ListJobsByGroupResponse {
	s.Headers = v
	return s
}

func (s *ListJobsByGroupResponse) SetStatusCode(v int32) *ListJobsByGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsByGroupResponse) SetBody(v *ListJobsByGroupResponseBody) *ListJobsByGroupResponse {
	s.Body = v
	return s
}

func (s *ListJobsByGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
