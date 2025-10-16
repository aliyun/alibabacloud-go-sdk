// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListAppInstanceGroupResponseBody) *ListAppInstanceGroupResponse
	GetBody() *ListAppInstanceGroupResponseBody
}

type ListAppInstanceGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppInstanceGroupResponse) GetBody() *ListAppInstanceGroupResponseBody {
	return s.Body
}

func (s *ListAppInstanceGroupResponse) SetHeaders(v map[string]*string) *ListAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstanceGroupResponse) SetStatusCode(v int32) *ListAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstanceGroupResponse) SetBody(v *ListAppInstanceGroupResponseBody) *ListAppInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *ListAppInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
