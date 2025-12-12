// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSessionGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSessionGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListSessionGroupResponseBody) *ListSessionGroupResponse
	GetBody() *ListSessionGroupResponseBody
}

type ListSessionGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSessionGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSessionGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSessionGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSessionGroupResponse) GetBody() *ListSessionGroupResponseBody {
	return s.Body
}

func (s *ListSessionGroupResponse) SetHeaders(v map[string]*string) *ListSessionGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSessionGroupResponse) SetStatusCode(v int32) *ListSessionGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionGroupResponse) SetBody(v *ListSessionGroupResponseBody) *ListSessionGroupResponse {
	s.Body = v
	return s
}

func (s *ListSessionGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
