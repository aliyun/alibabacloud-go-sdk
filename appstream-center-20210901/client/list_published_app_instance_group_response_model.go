// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAppInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishedAppInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishedAppInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishedAppInstanceGroupResponseBody) *ListPublishedAppInstanceGroupResponse
	GetBody() *ListPublishedAppInstanceGroupResponseBody
}

type ListPublishedAppInstanceGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedAppInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishedAppInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishedAppInstanceGroupResponse) GetBody() *ListPublishedAppInstanceGroupResponseBody {
	return s.Body
}

func (s *ListPublishedAppInstanceGroupResponse) SetHeaders(v map[string]*string) *ListPublishedAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedAppInstanceGroupResponse) SetStatusCode(v int32) *ListPublishedAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedAppInstanceGroupResponse) SetBody(v *ListPublishedAppInstanceGroupResponseBody) *ListPublishedAppInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *ListPublishedAppInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
