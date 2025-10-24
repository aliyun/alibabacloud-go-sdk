// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListTagGroupResponseBody) *ListTagGroupResponse
	GetBody() *ListTagGroupResponseBody
}

type ListTagGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagGroupResponse) GoString() string {
	return s.String()
}

func (s *ListTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagGroupResponse) GetBody() *ListTagGroupResponseBody {
	return s.Body
}

func (s *ListTagGroupResponse) SetHeaders(v map[string]*string) *ListTagGroupResponse {
	s.Headers = v
	return s
}

func (s *ListTagGroupResponse) SetStatusCode(v int32) *ListTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagGroupResponse) SetBody(v *ListTagGroupResponseBody) *ListTagGroupResponse {
	s.Body = v
	return s
}

func (s *ListTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
